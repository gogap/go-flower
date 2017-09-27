import "../aliyun"

include "../includes"


createProject = fn(clusterName, name, projectConf) {


	clients = new aliyun.AliyunClients(config)

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

	args = &aliyun_cs.ProjectCreationArgs{
		Name        : name,
		Description : projectConf.GetString("description"),
		Template    : string(tmplData),
		Version     : projectConf.GetString("version", "1.0.0"),
		LatestImage : projectConf.GetBoolean("latest-image", true),
	}

	err = projClient.CreateProject(args)

	if err!=nil {
		panic(err)
	}

	log.WithField("CODE", code).
		WithField("CLUSTER_NAME", clusterName).
		WithField("CLUSTER_ID", clusterId).
		WithField("PROJECT_NAME", name).
		Infoln("New project created")
}


main {

	log.WithField("CODE", code).Debug("Enter create_docker_project.ql")

	
	clustersConf = config.GetConfig("docker.clusters")

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
				panic("project config not set")
			}

			createProject(clusterName, projects[i], projectConf)
		}
	}
}