import "../aliyun"

include "../includes"

waitForClusterReady = fn(name) {

	clients = new aliyun.AliyunClients(config)

	clusterId, exist = ctx.Get("CLUSTERS_"+name)

	if !exist {

		clusters, err = clients.CS().DescribeClusters(name)


		if err!=nil {
			panic(err)
		}

		if len(clusters)==0 {
			panic("cluster of "+name+" not exist")
		}


		for i=0; i<len(clusters); i++{
			s = structs.new(clusters[i])
			if s.Field("Name").Value() == name {
				clusterId = s.Field("ClusterID").Value()
				exist = true
			}
		}
	}


	if !exist || len(clusterId) == 0 {
		panic("cluster of "+name+"'s id is empty")
	}


	log.WithField("CODE", code).
		WithField("CLUSTER_NAME", name).
		WithField("CLUSTER_ID", clusterId).
		Infoln("Wating for clutser ready")


	err = clients.CS().WaitForClusterAsyn(clusterId, aliyun_cs.Running, 600)

	if err !=nil {
		panic(err)
	}

	log.WithField("CODE", code).
		WithField("CLUSTER_NAME", name).
		WithField("CLUSTER_ID", clusterId).
		Infoln("Clutser is running")
}


main {

	log.WithField("CODE", code).Debug("Enter wait_for_cluster_ready.ql")

	
	clustersConf = config.GetConfig("docker.clusters")

	if clustersConf == nil {
		panic("clusters's config is empty")
	}

	clusters = clustersConf.Keys()

	if len(clusters) == 0 {
		panic("non cluster config exist")
	}

	wg = sync.NewWaitGroup()
	wg.Add(len(clusters))

	for i=0; i<len(clusters); i++{

		name = clusters[i]

		go fn(clusterName) {
			  defer wg.Done()
			  waitForClusterReady(clusterName)
		} (name)
	}

	wg.Wait()

	log.WithField("CODE", code).Infoln("All cluster are runnig")
}