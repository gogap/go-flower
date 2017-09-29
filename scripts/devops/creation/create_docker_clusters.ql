import "../aliyun"

include "../includes"



createCluster = fn(vpcId, vSwitchId, name, clusterConf) {


	clients = new aliyun.AliyunClients(CONFIG)

	clusters, err = clients.CS().DescribeClusters(name)

	if err!=nil {
		panic(err)
	}

	if len(clusters)>0 {

		for i=0; i<len(clusters); i++{
			s = structs.new(clusters[i])
			if s.Field("Name").Value() == name {
				clusterId = s.Field("ClusterID").Value()
				exist = true
			}
		}

		if exist {
			LOG.WithField("CODE", CODE).
				WithField("CLUSTER_NAME", name).
				WithField("CLUSTER_ID", clusterId).
				Infoln("Cluster already exist")

			return
		}
	}


	strCIDR = clusterConf.GetString("cluster.subnet-cidr")

	if len(strCIDR) == 0 {
		panic("cluster.subnet-cidr is empty")
	}


	args = &aliyun_cs.ClusterCreationArgs {
			Name             : name,
			Size             : clusterConf.GetInt64("cluster.size", 1),
			NetworkMode      : aliyun_cs.NetworkModeType(clusterConf.GetString("cluster.network-mode", "vpc")),
			SubnetCIDR       : clusterConf.GetString("cluster.subnet-cidr"),
			InstanceType     : clusterConf.GetString("cluster.instance-type", "ecs.n4.large"),
			VPCID            : vpcId,
			VSwitchID        : vSwitchId,
			Password		 : clusterConf.GetString("cluster.password"),
			DataDiskSize     : clusterConf.GetInt64("cluster.data-disk-size", 100),
			DataDiskCategory : aliyun_ecs.DiskCategory(clusterConf.GetString("cluster.data-disk-category")),
			ECSImageID       : clusterConf.GetString("cluster.ecs-image-id"),
			IOOptimized		 : aliyun_ecs.IoOptimized(clusterConf.GetString("cluster.io-optimized")),
	}

	resp, err = clients.CS().CreateCluster(clients.Region(), args)


	if err!=nil {
		panic(err)
	}

	// qlang's bug, 嵌入的结构貌似支持得不好, 所以用反射的方式来获取
	s = structs.new(resp)
	clusterID = s.Field("ClusterID").Value()


	CTX.Set("CLUSTERS_"+name, clusterID)

	LOG.WithField("CODE", CODE).
		WithField("CLUSTER_NAME", name).
		WithField("CLUSTER_ID", clusterID).
		Infoln("New cluster created")

	err = REDIS.Set(clusterKey, clusterID, 0).Err()

	if err!= nil {
		panic(err)
	}
}


main {


	LOG.WithField("CODE", CODE).Debug("Enter create_docker_clusters.ql")

	 if !ShouldExecute("docker_cluster") {
	 	return
	 }


	vpcId, vpcIdExist = CTX.Get("VPC_ID")

	if !vpcIdExist {
		panic("VPC_ID not in ctx")
	}


	vSwitchId, vSwitchIdExist = CTX.Get("VSWITCH_ID")

	if !vSwitchIdExist {
		panic("VSWITCH_ID not in ctx")
	}


	clustersConf = CONFIG.GetConfig("docker.clusters")

	if clustersConf == nil {
		panic("clusters's config is empty")
	}

	clusters = clustersConf.Keys()

	for i=0; i<len(clusters);i++{
		name = clusters[i]

		createCluster(vpcId, vSwitchId, name, clustersConf.GetConfig(name))
	}
}
