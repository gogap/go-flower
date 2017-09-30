
AliyunClients = class {

	fn _init(conf) {
		this.regionId = conf.GetString("region-id", "cn-beijing")
		this.akId = conf.GetString("access-key-id")
		this.akSecret = conf.GetString("access-key-secret")
		this.task = conf.GetString("task")


		this.CSClient=nil
		this.ECSClient=nil
		this.OSSClient=nil
		this.SLBClient=nil
		this.RDSClient=nil
	}

	fn CS() {
		if this.CSClient == nil {
			this.CSClient = aliyun_cs.client(this.akId, this.akSecret)
		}
		return this.CSClient
	}

	fn ECS() {
		if this.ECSClient == nil {
			this.ECSClient = aliyun_ecs.newClient(this.akId, this.akSecret)
		}
		return this.ECSClient
	}

	fn OSS() {
		if this.OSSClient == nil {
			this.OSSClient = aliyun_oss.newOSSClient(this.OSSRegion(), false, this.akId, this.akSecret, false)
		}
		return this.OSSClient
	}

	fn SLB() {
		if this.SLBClient == nil {
			this.SLBClient =  aliyun_slb.newClientWithRegion(
										aliyun_slb.SLBDefaultEndpoint, 
										this.akId, 
										this.akSecret, 
										this.OSSRegion())
		}
		
		return this.SLBClient
	}

	fn RDS() {
		if this.RDSClient == nil {
			this.RDSClient =  aliyun_rds.newRDSClient(
										this.akId, 
										this.akSecret, 
										this.OSSRegion())
		}
		
		return this.RDSClient
	}

	fn Region(region...) {
		regionId = this.regionId
		if len(region) >0 {
			regionId = region[0]
		}

		return aliyun_common.Region(regionId)
	}

	fn OSSRegion(region...) {
		regionId = this.regionId
		if len(region) >0 {
			regionId = region[0]
		}

		regionId = "oss-" + regionId

		return aliyun_oss.Region(regionId)
	}

	fn RegionId() {
		return this.regionId
	}

	fn OSSRegionId() {
		return "oss-" + this.regionId
	}
}

export AliyunClients