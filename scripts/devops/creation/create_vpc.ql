import "../aliyun"

include "../includes"

main {

	LOG.WithField("CODE", CODE).Debug("Enter create_vpc.ql")

	if !CanContinue("vpc") {
	 	return
	}

	vpcIdRedisKey = _CODE + ".devops.aliyun.ecs.vpc.id"

	// 检查是已经创建了这个VPC实例
	exists, err = _REDIS.Exists(vpcIdRedisKey).Result()

	if err!= nil {
		panic(err)
	}

	// 如果已经存在了，则就不再创建了
	if exists == 1  {

		vpcId = _REDIS.Get(vpcIdRedisKey).Val()

		CTX.Set("VPC_ID", vpcId)

		LOG.WithField("CODE", _CODE).WithField("VPC_ID", vpcId).Infoln("VPC already exist")

		return
	}


	clients = new aliyun.AliyunClients(_CONFIG)

	args = &aliyun_ecs.CreateVpcArgs {
		RegionId    : clients.Region(),
		VpcName     : _CONFIG.GetString("ecs.vpc.name", "vpc_"+_CODE),
		CidrBlock   : _CONFIG.GetString("ecs.vpc.cidr-block", "172.16.0.0/16"),
		Description : _CONFIG.GetString("ecs.vpc.description", "172.16.0.0/16"),
	}

	resp, err = clients.ECS().CreateVpc(args)

	if err!=nil {
	 	panic(err)
	}

	CTX.Set("VPC_ID", resp.VpcId)

	LOG.WithField("CODE", _CODE).WithField("VPC_ID", resp.VpcId).Infoln("New VPC created")


	err = _REDIS.Set(vpcIdRedisKey, resp.VpcId, 0).Err()

	if err!= nil {
		panic(err)
	}
}


