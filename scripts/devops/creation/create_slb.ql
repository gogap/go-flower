import "../aliyun"

include "../includes"


// TODO: Add create status to redis
createTCPListener = fn(slbId,slbName, name, conf) {

	args = &aliyun_slb.CreateLoadBalancerTCPListenerArgs {
					LoadBalancerId            : slbId,
					ListenerPort              : int(conf.GetInt64("listen-port")),
					BackendServerPort         : int(conf.GetInt64("server-port")),
					Bandwidth                 : int(conf.GetInt64("band-width")),
					Scheduler                 : aliyun_slb.SchedulerType(conf.GetString("scheduler","wrr")),
					PersistenceTimeout        : int(conf.GetInt64("persistence-timeout")),
					
					HealthCheck               : aliyun_slb.OnFlag,
					HealthCheckType           : aliyun_slb.HealthCheckType(conf.GetString("health-check.type","tcp")),
					HealthCheckDomain         : conf.GetString("health-check.domain"),
					HealthCheckURI            : conf.GetString("health-check.url"),
					HealthCheckConnectPort    : int(conf.GetInt64("health-check.connect-port")),
					HealthyThreshold          : int(conf.GetInt64("health-check.threshold", 3)),
					UnhealthyThreshold        : int(conf.GetInt64("health-check.unhealthy-threshold", 3)),
					HealthCheckConnectTimeout : int(conf.GetInt64("health-check.timeout", 5)),
					HealthCheckInterval       : int(conf.GetInt64("health-check.interval", 2)),
					HealthCheckHttpCode       : aliyun_slb.HealthCheckHttpCodeType(conf.GetString("health-check.http-code", "http_2xx")),

					VServerGroup              : aliyun_slb.FlagType(conf.GetString("vserver-group","on")),
					VServerGroupId            : conf.GetString("vserver-group-id"),
		}


	clients = new aliyun.AliyunClients(CONFIG)

	err = clients.SLB().CreateLoadBalancerTCPListener(args)

	if err!=nil {
		LOG.WithField("CODE", CODE).WithField("SLB_ID", slbId).WithField("SLB_NAME", slbName).WithField("SLB_NAME", name).Errorln(err)
		return
	}

	return
}

createHTTPListener = fn(slbId,slbName, name, conf) {

	
	args = &aliyun_slb.CreateLoadBalancerHTTPListenerArgs {
					LoadBalancerId            : slbId,
					ListenerPort              : int(conf.GetInt64("listen-port")),
					BackendServerPort         : int(conf.GetInt64("server-port")),
					Bandwidth                 : int(conf.GetInt64("band-width")),
					Scheduler                 : aliyun_slb.SchedulerType(conf.GetString("scheduler","wrr")),
					Gzip                      : aliyun_slb.FlagType(conf.GetString("gzip","on")),

					StickySession             : aliyun_slb.FlagType(conf.GetString("sticky-session", "off")),
					StickySessionType         : aliyun_slb.StickySessionType(conf.GetString("sticky-session-type", "insert")),
					
					CookieTimeout             : int(conf.GetInt64("cookie-timeout", 86400)),
					Cookie                    : conf.GetString("cookie"),

					HealthCheck               : aliyun_slb.FlagType(conf.GetString("health-check.check","on")),
					HealthCheckDomain         : conf.GetString("health-check.domain"),
					HealthCheckURI            : conf.GetString("health-check.url"),
					HealthCheckConnectPort    : int(conf.GetInt64("health-check.connect-port")),
					HealthyThreshold          : int(conf.GetInt64("health-check.threshold", 3)),
					UnhealthyThreshold        : int(conf.GetInt64("health-check.unhealthy-threshold", 3)),
					HealthCheckTimeout        : int(conf.GetInt64("health-check.timeout", 5)),
					HealthCheckInterval       : int(conf.GetInt64("health-check.interval", 2)),
					HealthCheckHttpCode       : aliyun_slb.HealthCheckHttpCodeType(conf.GetString("health-check.http-code", "http_2xx")),
					
					VServerGroup              : aliyun_slb.FlagType(conf.GetString("vserver-group","on")),
					VServerGroupId            : conf.GetString("vserver-group-id"),
		}


	clients = new aliyun.AliyunClients(CONFIG)

	err = clients.SLB().CreateLoadBalancerHTTPListener(args)

	if err!=nil {
		LOG.WithField("CODE", CODE).WithField("SLB_ID", slbId).WithField("SLB_NAME", name).Errorln(err)
		return
	}

	return

}


createSLBListeners = fn(slbId, slbName, conf) {

	clients = new aliyun.AliyunClients(CONFIG)

	tcpListenersConf = conf.GetConfig("tcp-listeners")

	if tcpListenersConf !=nil {

		tcpListeners = tcpListenersConf.Keys()

		for i=0;i<len(tcpListeners);i++ {
			
			listenerName = tcpListeners[i]

			listenerConf = tcpListenersConf.GetConfig(listenerName)

			createTCPListener(slbId, slbName, listenerName, listenerConf)
		}
	}

 

	httpListenersConf = conf.GetConfig("http-listeners")

	if httpListenersConf !=nil {

		httpListeners = httpListenersConf.Keys()

		for i=0;i<len(httpListeners);i++ {
			
			listenerName = httpListeners[i]

			listenerConf = httpListenersConf.GetConfig(listenerName)

			createHTTPListener(slbId, slbName, listenerName, listenerConf)
		}
	}

}

waitForSLBReady = fn(slbId,slbName) {

		clients = new aliyun.AliyunClients(CONFIG)

		describeArgs = &aliyun_slb.DescribeLoadBalancersArgs {
						RegionId            : clients.Region(),
						LoadBalancerId      : slbId,
  					}

		for {

  			LOG.WithField("CODE", CODE).WithField("SLB_ID", slbId).WithField("SLB_NAME", slbName).Debugln("Waiting for slb ready")

  			slbs, err = clients.SLB().DescribeLoadBalancers(describeArgs)

			if err!=nil {
				panic(err)
			}


			if len(slbs) == 0 {
				panic("describe slb failure, id:" + slbId)
			}

			slbResp = slbs[0]

			slbStatus = structs.new(slbResp).Field("LoadBalancerStatus").Value()

			if slbStatus == "active" {
				break
			}

			LOG.WithField("CODE", CODE).WithField("SLB_ID", slbId).WithField("SLB_NAME", slbName).WithField("STATUS", slbStatus).Debugln("SLB not ready")

			time.Sleep(time.Second*5)
  	}
}


createSLB = fn(vSwitchId, name, conf) {


		clients = new aliyun.AliyunClients(CONFIG)

	addressType = conf.GetString("address-type", "internet")
	chargeType = conf.GetString("charge-type", "paybytraffic")

	describeArgs = &aliyun_slb.DescribeLoadBalancersArgs {
						RegionId            : clients.Region(),
						LoadBalancerName    : name,
						VSwitchId           : vSwitchId,
  					}


  	slbs, err = clients.SLB().DescribeLoadBalancers(describeArgs)

  	if err!=nil {
  		panic(err)
  	}

 	if len(slbs)>0 {
 		LOG.WithField("CODE", CODE).WithField("SLB_NAME", name).Warnln("the same name slb already exist")
 		return
 	}

	args = &aliyun_slb.CreateLoadBalancerArgs {
						RegionId           : clients.Region(),
						LoadBalancerName   : name,
						AddressType        : aliyun_slb.AddressType(addressType),
						VSwitchId          : vSwitchId,
						InternetChargeType : aliyun_slb.InternetChargeType(chargeType),
						Bandwidth          : int(conf.GetInt64("band-width", 100)),
					}

	resp,err = clients.SLB().CreateLoadBalancer(args)

	if err!=nil {
		panic(err)
	}


	s = structs.new(resp)


	slbId = s.Field("LoadBalancerId").Value()

	LOG.WithField("CODE", CODE).WithField("SLB_ID", slbId).WithField("SLB_NAME", name).Infoln("New slb created")

	waitForSLBReady(slbId, name)

	createSLBListeners(slbId, name, conf)
}

main {

	LOG.WithField("CODE", CODE).Debug("Enter create_slb.ql")

	if !ShouldExecute("slb") {
	 	return
	}

	vSwitchId, exists = GetENV("VSWITCH_ID")

	if !exists {
		panic("vSwitchId not exist")
	}


	slbsConf = CONFIG.GetConfig("slb")

	if slbsConf == nil {
		return
	}

	sbls = slbsConf.Keys()

	if len(sbls) == 0 {
		return
	}

	for i=0;i<len(sbls);i++ {
		
		slbName = sbls[i]
		
		slbConf = slbsConf.GetConfig(slbName)

		createSLB(vSwitchId,slbName, slbConf)
	}

}