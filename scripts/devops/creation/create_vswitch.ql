import "../aliyun"

include "../includes"

main {

	log.WithField("CODE", code).Debug("Enter create_vswitch.ql")

	vpcId, vpcIdExist = ctx.Get("VPC_ID")


	// VPCID 不存在，说明未走VPC创建的流程
	if !vpcIdExist {
		panic("VPC_ID not in ctx")
	}

	vSwitchRedisKey = code + ".devops.aliyun.ecs.vswitch.id"

	// 检查是已经创建了这个VSwitch实例
	exists, err = redis.Exists(vSwitchRedisKey).Result()

	if err!= nil {
		panic(err)
	}

	// 如果已经存在了，则就不再创建了
	if exists == 1  {

		vSwitchId = redis.Get(vSwitchRedisKey).Val()

		ctx.Set("VSWITCH_ID", vSwitchId)

		log.WithField("CODE", code).WithField("VSWITCH_ID", vSwitchId).Infoln("VSwitch already exist")

		return
	}


	clients = new aliyun.AliyunClients(config)

	zoneId = config.GetString("ecs.vswitch.zone-id")

	if len(zoneId) > 0 {
		if !strings.HasPrefix(zoneId, clients.RegionId()) {
			panic("zone-id is illegal")
		}
	}

	args = &aliyun_ecs.CreateVSwitchArgs {
		VpcId       : vpcId,
		ZoneId      : config.GetString("ecs.vswitch.zone-id"),
		VSwitchName : config.GetString("ecs.vswitch.name", "vswitch_"+code),
		CidrBlock   : config.GetString("ecs.vswitch.cidr-block", "172.16.1.0/16"),
		Description : config.GetString("ecs.vswitch.description"),
	}

	vSwitchId, err = clients.ECS().CreateVSwitch(args)

	if err!=nil {
	 	panic(err)
	}

	ctx.Set("VSWITCH_ID", vSwitchId)

	log.WithField("CODE", code).WithField("VSWITCH_ID", vSwitchId).Infoln("New VSwitch created")

	err = redis.Set(vSwitchRedisKey, vSwitchId, 0).Err()

	if err!= nil {
		panic(err)
	}
}
