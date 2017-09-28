import "../aliyun"

include "../includes"


createBuckets = fn(name, perm) {


	clients = new aliyun.AliyunClients(config)

	bucket = clients.OSS().Bucket(name)

	binfo, err = bucket.Info()


	if err!=nil {
		if strings.Contains(err.Error(), "NoSuchBucket")  {
			err = nil
		} else {
			panic(err)
		}
	} else {
		LOG.WithField("CODE", CODE).WithField("BUCKET", name).Infoln("Bucket already exist")
		return
	}

	err = bucket.PutBucket(aliyun_oss.ACL(perm))

	if err!=nil {
		panic(err)
	}

	LOG.WithField("CODE", CODE).WithField("BUCKET", name).Infoln("New bucket created")
}

main {

	LOG.WithField("CODE", CODE).Debug("Enter create_oss_buckets.ql")

	if !CanContinue("oss_buckets") {
	 	return
	}

	ossConf = _CONFIG.GetConfig("oss")

	if ossConf == nil {
		panic("oss config not set")
	}

	buckets = ossConf.Keys()

	if len(buckets) == 0 {
		panic("oss bucket config is empty")
	}

	for i=0;i<len(buckets);i++ {
		bucket = buckets[i]
		bucketConf = ossConf.GetConfig(bucket)

		perm = bucketConf.GetString("perm", "private")

		createBuckets(bucket, perm)
	}

}