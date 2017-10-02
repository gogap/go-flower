import "../aliyun"
import "../shell"

include "../includes"

waitFor = fn (projClient,clusterName,clusterId,currentPorj, waitConf) {

	if waitConf == nil {
		return
	}

	waitProjects = waitConf.GetStringList("projects")

	println(waitProjects)

	if len(waitProjects) == 0 {
		return
	}

	wg = sync.NewWaitGroup()
	wg.Add(len(waitProjects))

	


	for i=0;i<len(waitProjects);i++{

		projName = waitProjects[i]

		go fn(projName){

			println(projName)

			defer wg.Done()
			
			resp,err = projClient.GetProject(projName)

			LOG.WithField("CODE", CODE).
				WithField("CLUSTER_NAME", clusterName).
				WithField("CLUSTER_ID", clusterId).
				WithField("CURRENT_PROJECT", currentPorj).
				WithField("WAIT_FOR_PROJECT", projName).
				Infoln("Wating for project ready")
			
			if err!=nil {
				LOG.WithField("CODE", CODE).
					WithField("CLUSTER_NAME", clusterName).
					WithField("CLUSTER_ID", clusterId).
					WithField("CURRENT_PROJECT", currentPorj).
					WithField("WAIT_FOR_PROJECT", projName).
					Errorln(err)
				return
			}

			s = structs.new(resp)

			for {
				
				currentStatus = s.Field("CurrentState").Value()

				if  currentStatus != "running" {

					LOG.WithField("CODE", CODE).
						WithField("CLUSTER_NAME", clusterName).
						WithField("CLUSTER_ID", clusterId).
						WithField("CURRENT_PROJECT", currentPorj).
						WithField("WAIT_FOR_PROJECT", projName).
						WithField("PROJECT_STATUS", currentStatus).
						Debugln("Waiting...")

					time.sleep(time.Second*5)
					continue
				}

				services = s.Field("Services").Value()

				total = len(services)
				for i=0;i<len(services);i++{

					srv = services[i]
					srvStruct  = structs.new(srv)

					if srvStruct.Field("CurrentState").Value() != "running" {
						LOG.WithField("CODE", CODE).
							WithField("CLUSTER_NAME", clusterName).
							WithField("CLUSTER_ID", clusterId).
							WithField("CURRENT_PROJECT", currentPorj).
							WithField("WAIT_FOR_PROJECT", projName).
							WithField("SERVICE_STATUS", services[i].CurrentState).
							Debugln("Waiting...")

						break
					} else {
						total=total-1
					}
				}

				if total>0{
					time.sleep(time.Second*5)
					continue
				} 

				break
			}
		}(projName)
	}

	wg.Wait()
	
}


createProject = fn(clusterName, name, projectConf) {


	clients = new aliyun.AliyunClients(CONFIG)
	
	clusterId = ""
	
	exist = false
	
	clusters, err = clients.CS().DescribeClusters(clusterName)


	if err!=nil {
		panic(err)
	}

	if len(clusters) == 0 {
		panic("cluster of "+clusterName+" not exist")
	}


	for i=0; i<len(clusters); i++{
		s = structs.new(clusters[i])
		if s.Field("Name").Value() == clusterName {
			clusterId = s.Field("ClusterID").Value()
			exist = true
		}
	}

	if !exist {
		panic("cluster "+clusterName + " not exist")
	}

	projClient, err = clients.CS().GetProjectClient(clusterId)

	if err!=nil {
		panic(err)
	}

	templateFile = projectConf.GetString("template")

	if len(templateFile) == 0 {
		panic("template-file not set")
	}

	tmplData, err = ioutil.readFile(templateFile)

	if err!=nil {
		panic(err)
	}

	if len(tmplData) == 0 {
		panic("template data is empty: "+ templateFile)
	}



	mapENVs = make(map[string]string)


	envsConf = projectConf.GetConfig("environment")

	waitConf = projectConf.GetConfig("wait")

	waitFor(projClient,clusterName, clusterId, name, waitConf)

	if envsConf!=nil {
			envKeys = envsConf.Keys()

			for i=0;i<len(envKeys);i++ {


				key = envKeys[i]
				val = envsConf.GetString(key)

				if strings.HasPrefix(val, "$.REDIS.") {
 
					redisKey = strings.TrimPrefix(val, "$.REDIS.")
		 
		 			val, vExist = GetENV(redisKey)

		 			if !vExist {
		 				panic("env not exist:" + redisKey)
		 			}

					if err!=nil {
						panic(err)
					}

				} elif strings.HasPrefix(val, "$.PWGEN.")  {
					
					strLen = strings.TrimPrefix(val, "$.PWGEN.")
					genLen,err = strconv.ParseInt(strLen,10,64)
					
					if err!=nil{
						genLen=16
						err=nil
					}


					ks = ["DEVOPS", clusterName, name, key]
 
					redisKey = strings.Join(ks, "_")
 
					redisKey = strings.ToUpper(redisKey)

					val, vExist = GetENV(redisKey)

					mapENVs[key] = val

					if vExist {
						continue
					}

					val = pwgen.alphaNum(genLen)

					err = SetENV(redisKey, val)

					if err!=nil {
						err = nil

						LOG.WithField("CODE", CODE).
							WithField("CLUSTER_NAME", clusterName).
							WithField("CLUSTER_ID", clusterId).
							WithField("PROJECT_NAME", name).
							WithField("PWGEN", val).
							Warnln("Set redis failure, so print this password")
					}
				} elif val == "$.INPUT" {
					println("xxxx")
		 			val = shell.Scanner.Input(key)
				}

				mapENVs[key] = val
			}
	}

	proj, err = projClient.GetProject(name)

	if err!=nil {
		if strings.Contains(err.Error(), "Status Code: 404") {
			err = nil
		} else {
			panic(err)
		}
	} else {
		LOG.WithField("CODE", CODE).
			WithField("CLUSTER_NAME", clusterName).
			WithField("CLUSTER_ID", clusterId).
			WithField("PROJECT_NAME", name).
			Infoln("Project already exist")
		return
	}


	args = &aliyun_cs.ProjectCreationArgs{
		Name        : name,
		Description : projectConf.GetString("description"),
		Template    : string(tmplData),
		Version     : projectConf.GetString("version", "1.0.0"),
		LatestImage : projectConf.GetBoolean("latest-image", true),
		Environment : mapENVs,
	}


	err = projClient.CreateProject(args)


	if err!=nil {
		panic(err)
	}

	LOG.WithField("CODE", CODE).
		WithField("CLUSTER_NAME", clusterName).
		WithField("CLUSTER_ID", clusterId).
		WithField("PROJECT_NAME", name).
		Infoln("New project created")
}


main {

	LOG.WithField("CODE", CODE).Debug("Enter create_docker_project.ql")

	if !ShouldExecute("docker_projects") {
	 	return
	 }
	
	clustersConf = CONFIG.GetConfig("docker.clusters")

	if clustersConf == nil {
		return
	}

	clusters = clustersConf.Keys()

	if len(clusters) == 0 {
		return
	}



	for i=0; i<len(clusters); i++{

		clusterName = clusters[i]

		clusterConf = clustersConf.GetConfig(clusterName)

		if clusterConf == nil {
			continue
		}

		projectsConf = clusterConf.GetConfig("projects")

		if projectsConf == nil {
			continue
		}

		projects = projectsConf.Keys()

		for i=0;i<len(projects);i++ {
			projectConf = projectsConf.GetConfig(projects[i])

			if projectConf == nil{
				panic("project CONFIG not set")
			}

			createProject(clusterName, projects[i], projectConf)
		}
	}
}