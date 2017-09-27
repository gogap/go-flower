import "../aliyun"

include "../includes"

main {

	log.WithField("CODE", code).Debug("Enter wait_for_ready.ql")

	vpcId, vpcIdExist = ctx.Get("VPC_ID")


	// VPCID 不存在，说明未走VPC创建的流程
	if !vpcIdExist {
		panic("VPC_ID not in ctx")
	}

	clients = new aliyun.AliyunClients(config)


	// 等待VPC进入到就绪状态

	log.WithField("CODE", code).WithField("VPC_ID", vpcId).Infoln("Wating for VPC ready")

	err = clients.ECS().WaitForVpcAvailable(clients.Region(), vpcId, 60)

	if err!=nil {
		panic(err)
	}

	log.WithField("CODE", code).WithField("VPC_ID", vpcId).Infoln("VPC current are available")
}