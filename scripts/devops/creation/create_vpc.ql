import "../aliyun"

include "../includes"

main {

	LOG.WithField("CODE", CODE).Debug("Enter create_vpc.ql")

	vpcId, exists = GetENV("VPC_ID")

	// 如果已经存在了，则就不再创建了
	if exists {

		LOG.WithField("CODE", CODE).WithField("VPC_ID", vpcId).Infoln("VPC already exist")

		return
	}


	if !ShouldExecute("vpc") {
	 	return
	}



	clients = new aliyun.AliyunClients(CONFIG)

	args = &aliyun_ecs.CreateVpcArgs {
		RegionId    : clients.Region(),
		VpcName     : CONFIG.GetString("ecs.vpc.name", "vpc_"+CODE),
		CidrBlock   : CONFIG.GetString("ecs.vpc.cidr-block", "172.16.0.0/16"),
		Description : CONFIG.GetString("ecs.vpc.description", "172.16.0.0/16"),
	}

	resp, err = clients.ECS().CreateVpc(args)

	if err!=nil {
	 	panic(err)
	}


	err = SetENV("VPC_ID", resp.VpcId) 
	
	if err !=nil {
		panic(err)
	}
 

	LOG.WithField("CODE", CODE).WithField("VPC_ID", resp.VpcId).Infoln("New VPC created")
	
}


