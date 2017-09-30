import "../aliyun"

include "../includes"

main {

	LOG.WithField("CODE", CODE).Debug("Enter create_vswitch.ql")


	vpcId, vpcIdExist = GetENV("VPC_ID")


	// VPCID 不存在，说明未走VPC创建的流程
	if !vpcIdExist {
		panic("VPC_ID not in env")
	}

	// 检查是已经创建了这个VSwitch实例
	vSwitchId, exists = GetENV("VSWITCH_ID")

	// 如果已经存在了，则就不再创建了
	if exists {
		LOG.WithField("CODE", CODE).WithField("VSWITCH_ID", vSwitchId).Infoln("VSwitch already exist")
		return
	}

	if !ShouldExecute("vswitch") {
	 	return
	}


	clients = new aliyun.AliyunClients(CONFIG)

	zoneId = CONFIG.GetString("ecs.vswitch.zone-id", "")

	if len(zoneId) > 0 {
		if !strings.HasPrefix(zoneId, clients.RegionId()) {
			panic("zone-id is illegal")
		}
	}

	args = &aliyun_ecs.CreateVSwitchArgs {
		VpcId       : vpcId,
		ZoneId      : CONFIG.GetString("ecs.vswitch.zone-id"),
		VSwitchName : CONFIG.GetString("ecs.vswitch.name", "vswitch_"+CODE),
		CidrBlock   : CONFIG.GetString("ecs.vswitch.cidr-block", "172.16.1.0/16"),
		Description : CONFIG.GetString("ecs.vswitch.description"),
	}

	vSwitchId, err = clients.ECS().CreateVSwitch(args)


	err = SetENV("VSWITCH_ID", vSwitchId)
	if  err!=nil {
		panic(err)
	}

	err = SetENV("VSWITCH_ID", vSwitchId)

	if err!=nil {
		panic(err)
	}

	LOG.WithField("CODE", CODE).WithField("VSWITCH_ID", vSwitchId).Infoln("New VSwitch created")
}
