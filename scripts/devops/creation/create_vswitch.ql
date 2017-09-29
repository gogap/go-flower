import "../aliyun"

include "../includes"

main {

	LOG.WithField("CODE", CODE).Debug("Enter create_vswitch.ql")

	if !ShouldExecute("vswitch") {
	 	return
	}

	vpcId, vpcIdExist = CTX.Get("VPC_ID")


	// VPCID 不存在，说明未走VPC创建的流程
	if !vpcIdExist {
		panic("VPC_ID not in ctx")
	}

	vSwitchRedisKey = CODE + ".devops.aliyun.ecs.vswitch.id"

	// 检查是已经创建了这个VSwitch实例
	exists, err = REDIS.Exists(vSwitchRedisKey).Result()

	if err!= nil {
		panic(err)
	}

	// 如果已经存在了，则就不再创建了
	if exists == 1  {

		vSwitchId = REDIS.Get(vSwitchRedisKey).Val()

		CTX.Set("VSWITCH_ID", vSwitchId)

		LOG.WithField("CODE", CODE).WithField("VSWITCH_ID", vSwitchId).Infoln("VSwitch already exist")

		return
	}


	clients = new aliyun.AliyunClients(config)

	zoneId = CONFIG.GetString("ecs.vswitch.zone-id")

	if len(zoneId) > 0 {
		if !strings.HasPrefix(zoneId, clients.RegionId()) {
			panic("zone-id is illegal")
		}
	}

	args = &aliyun_ecs.CreateVSwitchArgs {
		VpcId       : vpcId,
		ZoneId      : CONFIG.GetString("ecs.vswitch.zone-id"),
		VSwitchName : CONFIG.GetString("ecs.vswitch.name", "vswitch_"+CODE),
		CidrBlock   : CONFIG.GetString("ecs.vswitch.cidr-block", "172.16.1.0/16"),
		Description : CONFIG.GetString("ecs.vswitch.description"),
	}

	vSwitchId, err = clients.ECS().CreateVSwitch(args)

	if err!=nil {
	 	panic(err)
	}

	CTX.Set("VSWITCH_ID", vSwitchId)

	LOG.WithField("CODE", CODE).WithField("VSWITCH_ID", vSwitchId).Infoln("New VSwitch created")

	err = REDIS.Set(vSwitchRedisKey, vSwitchId, 0).Err()

	if err!= nil {
		panic(err)
	}
}
