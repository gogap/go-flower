package slb

import (
	"github.com/denverdino/aliyungo/slb"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/denverdino/aliyungo/slb",

	"ActiveStatus":           slb.ActiveStatus,
	"BackendServerPort":      slb.BackendServerPort,
	"Close":                  slb.Close,
	"Configuring":            slb.Configuring,
	"DefaultTimeout":         slb.DefaultTimeout,
	"DefaultWaitForInterval": slb.DefaultWaitForInterval,
	"HTTP":                    slb.HTTP,
	"HTTPHealthCheckType":     slb.HTTPHealthCheckType,
	"HTTPS":                   slb.HTTPS,
	"HTTP_2XX":                slb.HTTP_2XX,
	"HTTP_3XX":                slb.HTTP_3XX,
	"HTTP_4XX":                slb.HTTP_4XX,
	"HTTP_5XX":                slb.HTTP_5XX,
	"InactiveStatus":          slb.InactiveStatus,
	"InsertStickySessionType": slb.InsertStickySessionType,
	"InternetAddressType":     slb.InternetAddressType,
	"IntranetAddressType":     slb.IntranetAddressType,
	"LockedStatus":            slb.LockedStatus,
	"OffFlag":                 slb.OffFlag,
	"OnFlag":                  slb.OnFlag,
	"OpenWhileList":           slb.OpenWhileList,
	"PayByBandwidth":          slb.PayByBandwidth,
	"PayByTraffic":            slb.PayByTraffic,
	"Running":                 slb.Running,
	"SLBAPIVersion":           slb.SLBAPIVersion,
	"SLBDefaultEndpoint":      slb.SLBDefaultEndpoint,
	"SLBServiceCode":          slb.SLBServiceCode,
	"ServerStickySessionType": slb.ServerStickySessionType,
	"Starting":                slb.Starting,
	"Stopped":                 slb.Stopped,
	"Stopping":                slb.Stopping,
	"TCP":                     slb.TCP,
	"TCPHealthCheckType":      slb.TCPHealthCheckType,
	"UDP":                     slb.UDP,
	"WLCScheduler":            slb.WLCScheduler,
	"WRRScheduler":            slb.WRRScheduler,

	"InternetChargeType":      qlang.StructOf((*slb.InternetChargeType)(nil)),
	"AddressType":             qlang.StructOf((*slb.AddressType)(nil)),
	"SchedulerType":           qlang.StructOf((*slb.SchedulerType)(nil)),
	"FlagType":                qlang.StructOf((*slb.FlagType)(nil)),
	"HealthCheckType":         qlang.StructOf((*slb.HealthCheckType)(nil)),
	"HealthCheckHttpCodeType": qlang.StructOf((*slb.HealthCheckHttpCodeType)(nil)),
	"StickySessionType":       qlang.StructOf((*slb.StickySessionType)(nil)),

	"AddBackendServersArgs":                              qlang.StructOf((*slb.AddBackendServersArgs)(nil)),
	"AddBackendServersResponse":                          qlang.StructOf((*slb.AddBackendServersResponse)(nil)),
	"AddTagsArgs":                                        qlang.StructOf((*slb.AddTagsArgs)(nil)),
	"AddTagsResponse":                                    qlang.StructOf((*slb.AddTagsResponse)(nil)),
	"BackendServerType":                                  qlang.StructOf((*slb.BackendServerType)(nil)),
	"Client":                                             qlang.StructOf((*slb.Client)(nil)),
	"newClient":                                          slb.NewClient,
	"newClientWithEndpoint":                              slb.NewClientWithEndpoint,
	"newClientWithRegion":                                slb.NewClientWithRegion,
	"newSLBClient":                                       slb.NewSLBClient,
	"CommonListenerWhiteListItemArgs":                    qlang.StructOf((*slb.CommonListenerWhiteListItemArgs)(nil)),
	"CommonLoadBalancerListenerArgs":                     qlang.StructOf((*slb.CommonLoadBalancerListenerArgs)(nil)),
	"CommonLoadBalancerListenerResponse":                 qlang.StructOf((*slb.CommonLoadBalancerListenerResponse)(nil)),
	"CreateLoadBalancerArgs":                             qlang.StructOf((*slb.CreateLoadBalancerArgs)(nil)),
	"CreateLoadBalancerResponse":                         qlang.StructOf((*slb.CreateLoadBalancerResponse)(nil)),
	"CreateRulesArgs":                                    qlang.StructOf((*slb.CreateRulesArgs)(nil)),
	"CreateRulesResponse":                                qlang.StructOf((*slb.CreateRulesResponse)(nil)),
	"CreateVServerGroupArgs":                             qlang.StructOf((*slb.CreateVServerGroupArgs)(nil)),
	"CreateVServerGroupResponse":                         qlang.StructOf((*slb.CreateVServerGroupResponse)(nil)),
	"DeleteLoadBalancerArgs":                             qlang.StructOf((*slb.DeleteLoadBalancerArgs)(nil)),
	"DeleteLoadBalancerResponse":                         qlang.StructOf((*slb.DeleteLoadBalancerResponse)(nil)),
	"DeleteRulesArgs":                                    qlang.StructOf((*slb.DeleteRulesArgs)(nil)),
	"DeleteRulesResponse":                                qlang.StructOf((*slb.DeleteRulesResponse)(nil)),
	"DeleteServerCertificateArgs":                        qlang.StructOf((*slb.DeleteServerCertificateArgs)(nil)),
	"DeleteServerCertificateResponse":                    qlang.StructOf((*slb.DeleteServerCertificateResponse)(nil)),
	"DeleteVServerGroupArgs":                             qlang.StructOf((*slb.DeleteVServerGroupArgs)(nil)),
	"DeleteVServerGroupResponse":                         qlang.StructOf((*slb.DeleteVServerGroupResponse)(nil)),
	"DescribeHealthStatusArgs":                           qlang.StructOf((*slb.DescribeHealthStatusArgs)(nil)),
	"DescribeHealthStatusResponse":                       qlang.StructOf((*slb.DescribeHealthStatusResponse)(nil)),
	"DescribeListenerAccessControlAttributeResponse":     qlang.StructOf((*slb.DescribeListenerAccessControlAttributeResponse)(nil)),
	"DescribeLoadBalancerAttributeArgs":                  qlang.StructOf((*slb.DescribeLoadBalancerAttributeArgs)(nil)),
	"DescribeLoadBalancerAttributeResponse":              qlang.StructOf((*slb.DescribeLoadBalancerAttributeResponse)(nil)),
	"DescribeLoadBalancerHTTPListenerAttributeResponse":  qlang.StructOf((*slb.DescribeLoadBalancerHTTPListenerAttributeResponse)(nil)),
	"DescribeLoadBalancerHTTPSListenerAttributeResponse": qlang.StructOf((*slb.DescribeLoadBalancerHTTPSListenerAttributeResponse)(nil)),
	"DescribeLoadBalancerListenerAttributeResponse":      qlang.StructOf((*slb.DescribeLoadBalancerListenerAttributeResponse)(nil)),
	"DescribeLoadBalancerTCPListenerAttributeResponse":   qlang.StructOf((*slb.DescribeLoadBalancerTCPListenerAttributeResponse)(nil)),
	"DescribeLoadBalancerUDPListenerAttributeResponse":   qlang.StructOf((*slb.DescribeLoadBalancerUDPListenerAttributeResponse)(nil)),
	"DescribeLoadBalancersArgs":                          qlang.StructOf((*slb.DescribeLoadBalancersArgs)(nil)),
	"DescribeLoadBalancersResponse":                      qlang.StructOf((*slb.DescribeLoadBalancersResponse)(nil)),
	"DescribeRegionsArgs":                                qlang.StructOf((*slb.DescribeRegionsArgs)(nil)),
	"DescribeRegionsResponse":                            qlang.StructOf((*slb.DescribeRegionsResponse)(nil)),
	"DescribeRuleAttributeArgs":                          qlang.StructOf((*slb.DescribeRuleAttributeArgs)(nil)),
	"DescribeRuleAttributeResponse":                      qlang.StructOf((*slb.DescribeRuleAttributeResponse)(nil)),
	"DescribeRulesArgs":                                  qlang.StructOf((*slb.DescribeRulesArgs)(nil)),
	"DescribeRulesResponse":                              qlang.StructOf((*slb.DescribeRulesResponse)(nil)),
	"DescribeServerCertificatesArgs":                     qlang.StructOf((*slb.DescribeServerCertificatesArgs)(nil)),
	"DescribeServerCertificatesResponse":                 qlang.StructOf((*slb.DescribeServerCertificatesResponse)(nil)),
	"DescribeTagsArgs":                                   qlang.StructOf((*slb.DescribeTagsArgs)(nil)),
	"DescribeTagsResponse":                               qlang.StructOf((*slb.DescribeTagsResponse)(nil)),
	"DescribeVServerGroupAttributeArgs":                  qlang.StructOf((*slb.DescribeVServerGroupAttributeArgs)(nil)),
	"DescribeVServerGroupsArgs":                          qlang.StructOf((*slb.DescribeVServerGroupsArgs)(nil)),
	"DescribeVServerGroupsResponse":                      qlang.StructOf((*slb.DescribeVServerGroupsResponse)(nil)),
	"HTTPListenerType":                                   qlang.StructOf((*slb.HTTPListenerType)(nil)),
	"CreateLoadBalancerHTTPListenerArgs":                 qlang.StructOf((*slb.CreateLoadBalancerHTTPListenerArgs)(nil)),
	"HTTPSListenerType":                                  qlang.StructOf((*slb.HTTPSListenerType)(nil)),
	"HealthStatusType":                                   qlang.StructOf((*slb.HealthStatusType)(nil)),
	"ListenerPortAndProtocolType":                        qlang.StructOf((*slb.ListenerPortAndProtocolType)(nil)),
	"LoadBalancerType":                                   qlang.StructOf((*slb.LoadBalancerType)(nil)),
	"ModifyLoadBalancerInternetSpecArgs":                 qlang.StructOf((*slb.ModifyLoadBalancerInternetSpecArgs)(nil)),
	"ModifyLoadBalancerInternetSpecResponse":             qlang.StructOf((*slb.ModifyLoadBalancerInternetSpecResponse)(nil)),
	"ModifyVServerGroupBackendServersArgs":               qlang.StructOf((*slb.ModifyVServerGroupBackendServersArgs)(nil)),
	"RegionType":                                         qlang.StructOf((*slb.RegionType)(nil)),
	"RemoveBackendServersArgs":                           qlang.StructOf((*slb.RemoveBackendServersArgs)(nil)),
	"RemoveBackendServersResponse":                       qlang.StructOf((*slb.RemoveBackendServersResponse)(nil)),
	"RemoveTagsArgs":                                     qlang.StructOf((*slb.RemoveTagsArgs)(nil)),
	"RemoveTagsResponse":                                 qlang.StructOf((*slb.RemoveTagsResponse)(nil)),
	"Rule":                                               qlang.StructOf((*slb.Rule)(nil)),
	"ServerCertificateType":                              qlang.StructOf((*slb.ServerCertificateType)(nil)),
	"SetListenerAccessControlStatusArgs":                 qlang.StructOf((*slb.SetListenerAccessControlStatusArgs)(nil)),
	"SetLoadBalancerNameArgs":                            qlang.StructOf((*slb.SetLoadBalancerNameArgs)(nil)),
	"SetLoadBalancerNameResponse":                        qlang.StructOf((*slb.SetLoadBalancerNameResponse)(nil)),
	"SetLoadBalancerStatusArgs":                          qlang.StructOf((*slb.SetLoadBalancerStatusArgs)(nil)),
	"SetLoadBalancerStatusResponse":                      qlang.StructOf((*slb.SetLoadBalancerStatusResponse)(nil)),
	"SetRuleArgs":                                        qlang.StructOf((*slb.SetRuleArgs)(nil)),
	"SetRuleResponse":                                    qlang.StructOf((*slb.SetRuleResponse)(nil)),
	"SetServerCertificateNameArgs":                       qlang.StructOf((*slb.SetServerCertificateNameArgs)(nil)),
	"SetServerCertificateNameResponse":                   qlang.StructOf((*slb.SetServerCertificateNameResponse)(nil)),
	"SetVServerGroupAttributeArgs":                       qlang.StructOf((*slb.SetVServerGroupAttributeArgs)(nil)),
	"SetVServerGroupAttributeResponse":                   qlang.StructOf((*slb.SetVServerGroupAttributeResponse)(nil)),
	"TCPListenerType":                                    qlang.StructOf((*slb.TCPListenerType)(nil)),
	"CreateLoadBalancerTCPListenerArgs":                  qlang.StructOf((*slb.CreateLoadBalancerTCPListenerArgs)(nil)),
	"TagItem":                                            qlang.StructOf((*slb.TagItem)(nil)),
	"TagItemType":                                        qlang.StructOf((*slb.TagItemType)(nil)),
	"UDPListenerType":                                    qlang.StructOf((*slb.UDPListenerType)(nil)),
	"UploadServerCertificateArgs":                        qlang.StructOf((*slb.UploadServerCertificateArgs)(nil)),
	"UploadServerCertificateResponse":                    qlang.StructOf((*slb.UploadServerCertificateResponse)(nil)),
	"VBackendServerType":                                 qlang.StructOf((*slb.VBackendServerType)(nil)),
	"VBackendServers":                                    qlang.StructOf((*slb.VBackendServers)(nil)),
	"VServerGroup":                                       qlang.StructOf((*slb.VServerGroup)(nil)),
}
