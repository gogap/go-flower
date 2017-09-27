import "../aliyun"

include "../includes"

main {

	log.WithField("CODE", code).Debug("Enter create_vpc.ql")

	vpcIdRedisKey = code + ".devops.aliyun.ecs.vpc.id"

	// 检查是已经创建了这个VPC实例
	exists, err = redis.Exists(vpcIdRedisKey).Result()

	if err!= nil {
		panic(err)
	}

	// 如果已经存在了，则就不再创建了
	if exists == 1  {

		vpcId = redis.Get(vpcIdRedisKey).Val()

		ctx.Set("VPC_ID", vpcId)

		log.WithField("CODE", code).WithField("VPC_ID", vpcId).Infoln("VPC already exist")

		return
	}


	clients = new aliyun.AliyunClients(config)

	args = &aliyun_ecs.CreateVpcArgs {
		RegionId    : clients.Region(),
		VpcName     : config.GetString("ecs.vpc.name", "vpc_"+code),
		CidrBlock   : config.GetString("ecs.vpc.cidr-block", "172.16.0.0/16"),
		Description : config.GetString("ecs.vpc.description", "172.16.0.0/16"),
	}

	resp, err = clients.ECS().CreateVpc(args)

	if err!=nil {
	 	panic(err)
	}

	ctx.Set("VPC_ID", resp.VpcId)

	log.WithField("CODE", code).WithField("VPC_ID", resp.VpcId).Infoln("New VPC created")


	err = redis.Set(vpcIdRedisKey, resp.VpcId, 0).Err()

	if err!= nil {
		panic(err)
	}
}


