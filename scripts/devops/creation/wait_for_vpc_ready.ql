import "../aliyun"

include "../includes"

main {

	LOG.WithField("CODE", CODE).Debug("Enter wait_for_ready.ql")

	if !CanContinue("vpc") {
	 	return
	}

	vpcId, vpcIdExist = CTX.Get("VPC_ID")


	// VPCID 不存在，说明未走VPC创建的流程
	if !vpcIdExist {
		panic("VPC_ID not in ctx")
	}

	clients = new aliyun.AliyunClients(config)


	// 等待VPC进入到就绪状态

	LOG.WithField("CODE", CODE).WithField("VPC_ID", vpcId).Infoln("Wating for VPC ready")

	err = clients.ECS().WaitForVpcAvailable(clients.Region(), vpcId, 60)

	if err!=nil {
		panic(err)
	}

	LOG.WithField("CODE", CODE).WithField("VPC_ID", vpcId).Infoln("VPC current are available")
}