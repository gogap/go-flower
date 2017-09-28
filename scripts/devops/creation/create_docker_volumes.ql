import "../aliyun"

include "../includes"


createVolume = fn(clusterName, name, volumeConf) {

	driver = volumeConf.GetString("driver")

	if len(driver) == 0 {
		panic("volume driver not set")
	}

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

	args = &aliyun_cs.VolumeCreationArgs{
		Name : name,
		Driver:aliyun_cs.VolumeDriverType(driver),
	}

	if driver == "ossfs" {

		bucket = volumeConf.GetString("options.bucket")
		url = volumeConf.GetString("options.url")
		akId = volumeConf.GetString("options.ak-id")
		akSec = volumeConf.GetString("options.ak-secret")

		if len(bucket) == 0 {
			panic("bucket not set")
		} 

		if len(url) == 0 {
			panic("url not set")
		} 

		if len(akId) == 0 {
			panic("oss volume AccessKeyId is empty")
		}

		if len(akSec) == 0 {
			panic("oss volume AccessKeySecret is empty")
		}

		args.DriverOpts = &aliyun_cs.OSSOpts{
			Bucket          : bucket,
			AccessKeyId     : akId,
			AccessKeySecret : akSec,
			URL             : url,
			NoStatCache     : volumeConf.GetString("options.no_stat_cache", "false"),
			OtherOpts       : volumeConf.GetString("options.other_opts"),
		}

	} elif driver == "nas" {

		args.DriverOpts = &aliyun_cs.NASOpts{
			DiskId : volumeConf.GetString("options.disk-id"),
			Host   : volumeConf.GetString("options.host"),
			Path   : volumeConf.GetString("options.path"),
			Mode   : volumeConf.GetString("options.mode"),
		}
	}

	err = projClient.CreateVolume(args)

	if err!=nil {
		panic(err)
	}

	LOG.WithField("CODE", CODE).
		WithField("CLUSTER_NAME", clusterName).
		WithField("CLUSTER_ID", clusterId).
		WithField("VOLUME_NAME", name).
		Infoln("New data volume created")
}


main {

	LOG.WithField("CODE", CODE).Debug("Enter create_docker_volumes.ql")

	if !CanContinue("docker_volumes") {
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

		volumesConf = clusterConf.GetConfig("volumes")

		if volumesConf == nil {
			continue
		}

		volumes = volumesConf.Keys()

		for i=0;i<len(volumes);i++ {
			volumeConf = volumesConf.GetConfig(volumes[i])

			if volumeConf == nil{
				panic("volume config not set")
			}

			createVolume(clusterName, volumes[i], volumeConf)
		}
	}
}