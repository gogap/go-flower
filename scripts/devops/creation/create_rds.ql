import "../aliyun"

include "../includes"

createAccounts = fn (rdsName, instanceId, rdsConf) {

	LOG.WithField("CODE", CODE).
		WithField("RDS", rdsName).
		WithField("INSTANCE-ID", instanceId).
		Debugln("Begin create accounts")

	accountsConf = rdsConf.GetConfig("accounts")

	accounts = accountsConf.Keys()

	if accounts == nil {
		return
	}

	clients = new aliyun.AliyunClients(CONFIG)

	for i=0;i<len(accounts);i++{

		accountConf = accountsConf.GetConfig(accounts[i])


		descArgs =  &aliyun_rds.DescribeAccountsArgs {
									DBInstanceId : instanceId,
									AccountName  : accounts[i],
								}


		descResp, err = clients.RDS().DescribeAccounts(descArgs)

		if err!=nil {
			panic(err)
		}

		dbInsResps = structs.new(descResp.Accounts)
		dbInsRespsField = dbInsResps.Field("DBInstanceAccount").Value()

		needCreateAccount = true

		if len(dbInsRespsField) > 0 {

			LOG.WithField("CODE", CODE).
				WithField("RDS", rdsName).
				WithField("INSTANCE-ID", instanceId).
				WithField("ACCOUNT", accounts[i]).
				Debugln("Account already exist")

			needCreateAccount = false
		}


		if needCreateAccount {

				args = &aliyun_rds.CreateAccountArgs{
					DBInstanceId       : instanceId,
					AccountName        : accounts[i],
					AccountPassword    : accountConf.GetString("password"),
					AccountType        : aliyun_rds.AccountType(accountConf.GetString("type", "Normal")),
					AccountDescription : accountConf.GetString("description"),
				}


				_, err = clients.RDS().CreateAccount(args)

				if err!=nil {
					panic(err)
				}

				LOG.WithField("CODE", CODE).
					WithField("RDS", rdsName).
					WithField("INSTANCE-ID", instanceId).
					WithField("ACCOUNT", accounts[i]).
					Debugln("New account created")
		}

		accountPWDEnvKey = "DEVOPS_ALIYUN_RDS_"+strings.ToUpper(rdsName+"_PASS_"+ accounts[i]) 

		err = SetENV(accountPWDEnvKey, accountConf.GetString("password"))

		if err!=nil {
			panic(err)
		}

		err = clients.RDS().WaitForAccount(instanceId, accounts[i], aliyun_rds.Available, 300)

		if err!=nil {
			panic("Waiting for account available failure")	
		} 


		LOG.WithField("CODE", CODE).
				WithField("RDS", rdsName).
				WithField("INSTANCE-ID", instanceId).
				WithField("ACCOUNT", accounts[i]).
				Infoln("Account is available")

		accountDBsConf = accountConf.GetConfig("databases")

		if accountDBsConf == nil {
			continue
		}

		dbs = accountDBsConf.Keys()


		for j=0;j<len(dbs);j++ {

			dbPrivConf = accountDBsConf.GetConfig(dbs[j])

			grantArgs =  &aliyun_rds.GrantAccountPrivilegeArgs{
							DBInstanceId     : instanceId,
							AccountName      : accounts[i],
							DBName           : dbs[j],
							AccountPrivilege : aliyun_rds.AccountPrivilege(dbPrivConf.GetString("privilege", "ReadWrite")),
						}


			_, err = clients.RDS().GrantAccountPrivilege(grantArgs)

			if err!=nil {
				panic(err)
			}

			LOG.WithField("CODE", CODE).
				WithField("RDS", rdsName).
				WithField("INSTANCE-ID", instanceId).
				WithField("ACCOUNT", accounts[i]).
				WithField("DB", dbs[j]).
				Debugln("Grant account privilege success")
		}
	}
}


waitForAllDatabasesReady = fn(rdsName, instanceId, rdsConf) {

	LOG.WithField("CODE", CODE).
		WithField("RDS", rdsName).
		WithField("INSTANCE-ID", instanceId).
		Infoln("Waiting for rds instance's databases ready")

	dbsConf = rdsConf.GetConfig("databases")

	clients = new aliyun.AliyunClients(CONFIG)

	dbNames = dbsConf.Keys()

	err = clients.RDS().WaitForAllDatabase(instanceId, dbNames, aliyun_rds.Running, 600)

	if err!=nil {
		panic(err)
	}
}

createInstanceDatabases = fn(rdsName, instanceId, rdsConf) {

	LOG.WithField("CODE", CODE).
		WithField("RDS", rdsName).
		WithField("INSTANCE-ID", instanceId).
		Debugln("Begin create databases")

	dbsConf = rdsConf.GetConfig("databases")

	if dbsConf == nil {
		return
	}


	clients = new aliyun.AliyunClients(CONFIG)

	dbNames = dbsConf.Keys()


	for i=0;i<len(dbNames);i++{

		dbName = dbNames[i]

		dbConf = dbsConf.GetConfig(dbName)

		describDbArgs = &aliyun_rds.DescribeDatabasesArgs {
			DBInstanceId : instanceId,
			DBName       : dbName,
		}


		respDescib, err = clients.RDS().DescribeDatabases(describDbArgs)

		dbInsResps = structs.new(respDescib.Databases)
		dbInsRespsField = dbInsResps.Field("Database").Value()

		if len(dbInsRespsField) > 0 {
			LOG.WithField("CODE", CODE).
				WithField("RDS", rdsName).
				WithField("INSTANCE-ID", instanceId).
				WithField("DB", dbName).
				Infoln("Database already exist")

			continue
		}



		args = &aliyun_rds.CreateDatabaseArgs {
			DBInstanceId     : instanceId,
			DBName           : dbName,
			CharacterSetName : dbConf.GetString("character-set"),
			DBDescription    : dbConf.GetString("description"),
		}
		
		resp, err = clients.RDS().CreateDatabase(args)

		if err!=nil{

			LOG.WithField("CODE", CODE).
				WithField("RDS", rdsName).
				WithField("INSTANCE-ID", instanceId).
				WithField("DB", dbName).
				Errorln(err)

			err = nil
		}

		LOG.WithField("CODE", CODE).
			WithField("RDS", rdsName).
			WithField("INSTANCE-ID", instanceId).
			WithField("DB", dbName).
			Infoln("New database created")
	}
}




waitForRDSReady = fn(rdsName, instanceId) {


	LOG.WithField("CODE", CODE).
		WithField("RDS", rdsName).
		WithField("INSTANCE-ID", instanceId).
		Infoln("Waiting for rds instance ready")

	clients = new aliyun.AliyunClients(CONFIG)

	err = clients.RDS().WaitForInstance(instanceId, aliyun_rds.Running, 600)

	if err!=nil {
		LOG.WithField("CODE", CODE).
		WithField("RDS", rdsName).
		WithField("INSTANCE-ID", instanceId).
		Errorln(err)
		return
	}

	LOG.WithField("CODE", CODE).
		WithField("RDS", rdsName).
		WithField("INSTANCE-ID", instanceId).
		Infoln("RDS instance is ready")
}


createRDS = fn(vpcId, vSwitchId, rdsName, rdsConf) {

	rdsKey = CODE + ".devops.aliyun.rds."+ rdsName

	exists, err = REDIS.Exists(rdsKey).Result()

	if err!= nil {
		panic(err)
	}


	if exists == 1  {

		kvs, err = REDIS.HGetAll(rdsKey).Result()

		if err!=nil{
			panic(err)
		}

		for k,v = range kvs {

			if k == "_" {
				continue
			}

			envKey =  "DEVOPS_ALIYUN_RDS_"+ strings.ToUpper(rdsName  + "_" + k)

			err = SetENV(envKey, v)

			if err!=nil{
				panic(err)
			}
			
			LOG.WithField("CODE", CODE).
				WithField("ENV", envKey).
				Debugln("Env setted")
		}

		LOG.WithField("CODE", CODE).
			WithField("RDS", rdsName).
			WithField("INSTANCE-ID", kvs["id"]).
			Infoln("RDS instance already created")

		return kvs["id"]
	}

	clients = new aliyun.AliyunClients(CONFIG)

	args = &aliyun_rds.CreateDBInstanceArgs {
				RegionId              : clients.Region(),
				Engine                : aliyun_rds.Engine(rdsConf.GetString("engine","MySQL")),
				EngineVersion         : rdsConf.GetString("engine-version","5.6"),
				PayType               : aliyun_rds.DBPayType(rdsConf.GetString("pay-type","Postpaid")),

				DBInstanceClass       : rdsConf.GetString("instance-class","rds.mys2.small"),
				DBInstanceStorage     : int(rdsConf.GetInt64("instance-storage",5)),
				DBInstanceNetType     : aliyun_common.NetType(rdsConf.GetString("instance-net-type","Internet")),
				DBInstanceDescription : rdsConf.GetString("instance-description"),
				InstanceNetworkType   : aliyun_common.NetworkType(rdsConf.GetString("instance-network-type","VPC")),

				VPCId                 : vpcId,
				VSwitchId             : vSwitchId,
				ZoneId                : rdsConf.GetString("zone-id"),

				UsedTime              : rdsConf.GetString("used-time"),
				Period                : aliyun_common.TimeType(rdsConf.GetString("period")),
 

				ConnectionMode        : aliyun_rds.ConnectionMode(rdsConf.GetString("connection-mode","Performance")),
				SecurityIPList        : rdsConf.GetString("security-ip-list", "172.18.0.0/24"),
				PrivateIpAddress      : rdsConf.GetString("private-ip-address", ""),

				ClientToken           : "",
	}

	resp,err = clients.RDS().CreateDBInstance(args)
 
	if err!=nil {
		panic(err)
	}


	s = structs.new(resp)

	instId  = s.Field("DBInstanceId").Value()
	connStr = s.Field("ConnectionString").Value()
	port    = s.Field("Port").Value()

	LOG.WithField("CODE", CODE).
		WithField("RDS", rdsName).
		WithField("INSTANCE-ID", instId).
		WithField("CONN", connStr).
		WithField("PORT", port).
		Infoln("New RDS created")

	
	values = {"id": instId, "conn": connStr, "port": port, "_": true} // for make map[string]interface{}

	err = REDIS.HMSet(rdsKey, values).Err()

	if err!= nil {
		panic(err)
	}



	for k,v = range values {

		if k == "_" {
			continue
		}

		envKey = "DEVOPS_ALIYUN_RDS_"+ strings.ToUpper(rdsName  + "_" + k)

		err = SetENV(envKey, v)

		if err!=nil{
			panic(err)
		}
	}

	return instId
}

main {

	LOG.WithField("CODE", CODE).Debug("Enter create_rds.ql")

	if !ShouldExecute("rds") {
	 	return
	}


	vpcId, vpcIdExist = GetENV("VPC_ID")

	if !vpcIdExist {
		panic("VPC_ID not in env")
	}


	vSwitchId, vSwitchIdExist = GetENV("VSWITCH_ID")

	if !vSwitchIdExist {
		panic("VSWITCH_ID not in env")
	}


	rdssConf = CONFIG.GetConfig("rds")

	if rdssConf == nil {
		panic("rds config not set")
	}

	RDSs = rdssConf.Keys()

	if len(RDSs) == 0 {
		panic("rds config is empty")
	}

	wg = sync.NewWaitGroup()

	instanceIds = make([]string,len(RDSs),len(RDSs))

 
	for i=0;i<len(RDSs);i++ {
		
		rdsName = RDSs[i]
		rdsConf = rdssConf.GetConfig(rdsName)
 
		wg.Add(1)

		instId = createRDS(vpcId, vSwitchId, rdsName, rdsConf)


		instanceIds[i] = instId



		if len(instId) > 0 {
			go fn(rdsName, instId, rdsConf) {
				 defer wg.Done()

				 waitForRDSReady(rdsName, instId)

			} (rdsName, instId, rdsConf)
		}
	}

	wg.Wait()

	LOG.WithField("CODE", CODE).
		Infoln("All RDS instance are ready")



	for i=0;i<len(RDSs);i++ {

		id = instanceIds[i]
		rdsName = RDSs[i]

		rdsConf = rdssConf.GetConfig(rdsName)

		createInstanceDatabases(rdsName, id, rdsConf)
	}



	wg2 = sync.NewWaitGroup()
	wg2.Add(len(instanceIds))

	for i=0;i<len(RDSs);i++ {

		id = instanceIds[i]
		rdsName = RDSs[i]

		rdsConf = rdssConf.GetConfig(rdsName)

		go fn(rdsName, id, rdsConf){
			defer wg2.Done()
			waitForAllDatabasesReady(rdsName, id, rdsConf)
		}(rdsName, id, rdsConf)
	}

	wg2.Wait()

	for i=0;i<len(RDSs);i++ {

		id = instanceIds[i]
		rdsName = RDSs[i]

		rdsConf = rdssConf.GetConfig(rdsName)

		createAccounts(rdsName, id, rdsConf)
	}



	LOG.WithField("CODE", CODE).
		Infoln("All RDS instances, databases and accounts are ready")
	
}