package dns

import (
	"github.com/denverdino/aliyungo/dns"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/denverdino/aliyungo/dns",

	"AAAARecord":            dns.AAAARecord,
	"ARecord":               dns.ARecord,
	"CNAMERecord":           dns.CNAMERecord,
	"DNSAPIVersion":         dns.DNSAPIVersion,
	"DNSDefaultEndpoint":    dns.DNSDefaultEndpoint,
	"DNSDefaultEndpointNew": dns.DNSDefaultEndpointNew,
	"ForwordURLRecord":      dns.ForwordURLRecord,
	"MXRecord":              dns.MXRecord,
	"NSRecord":              dns.NSRecord,
	"RedirectURLRecord":     dns.RedirectURLRecord,
	"SRVRecord":             dns.SRVRecord,
	"TXTRecord":             dns.TXTRecord,

	"AddDomainArgs":                       qlang.StructOf((*dns.AddDomainArgs)(nil)),
	"AddDomainGroupArgs":                  qlang.StructOf((*dns.AddDomainGroupArgs)(nil)),
	"AddDomainGroupResponse":              qlang.StructOf((*dns.AddDomainGroupResponse)(nil)),
	"AddDomainRecordArgs":                 qlang.StructOf((*dns.AddDomainRecordArgs)(nil)),
	"AddDomainRecordResponse":             qlang.StructOf((*dns.AddDomainRecordResponse)(nil)),
	"AddDomainResponse":                   qlang.StructOf((*dns.AddDomainResponse)(nil)),
	"ChangeDomainGroupArgs":               qlang.StructOf((*dns.ChangeDomainGroupArgs)(nil)),
	"ChangeDomainGroupResponse":           qlang.StructOf((*dns.ChangeDomainGroupResponse)(nil)),
	"Client":                              qlang.StructOf((*dns.Client)(nil)),
	"newClient":                           dns.NewClient,
	"newClientNew":                        dns.NewClientNew,
	"newClientWithEndpoint":               dns.NewClientWithEndpoint,
	"newCustomClient":                     dns.NewCustomClient,
	"DeleteDomainArgs":                    qlang.StructOf((*dns.DeleteDomainArgs)(nil)),
	"DeleteDomainGroupArgs":               qlang.StructOf((*dns.DeleteDomainGroupArgs)(nil)),
	"DeleteDomainGroupResponse":           qlang.StructOf((*dns.DeleteDomainGroupResponse)(nil)),
	"DeleteDomainRecordArgs":              qlang.StructOf((*dns.DeleteDomainRecordArgs)(nil)),
	"DeleteDomainRecordResponse":          qlang.StructOf((*dns.DeleteDomainRecordResponse)(nil)),
	"DeleteDomainResponse":                qlang.StructOf((*dns.DeleteDomainResponse)(nil)),
	"DeleteSubDomainRecordsArgs":          qlang.StructOf((*dns.DeleteSubDomainRecordsArgs)(nil)),
	"DeleteSubDomainRecordsResponse":      qlang.StructOf((*dns.DeleteSubDomainRecordsResponse)(nil)),
	"DescribeDomainGroupsArgs":            qlang.StructOf((*dns.DescribeDomainGroupsArgs)(nil)),
	"DescribeDomainInfoArgs":              qlang.StructOf((*dns.DescribeDomainInfoArgs)(nil)),
	"DescribeDomainRecordInfoArgs":        qlang.StructOf((*dns.DescribeDomainRecordInfoArgs)(nil)),
	"DescribeDomainRecordInfoNewArgs":     qlang.StructOf((*dns.DescribeDomainRecordInfoNewArgs)(nil)),
	"DescribeDomainRecordInfoNewResponse": qlang.StructOf((*dns.DescribeDomainRecordInfoNewResponse)(nil)),
	"DescribeDomainRecordInfoResponse":    qlang.StructOf((*dns.DescribeDomainRecordInfoResponse)(nil)),
	"DescribeDomainRecordsArgs":           qlang.StructOf((*dns.DescribeDomainRecordsArgs)(nil)),
	"DescribeDomainRecordsNewArgs":        qlang.StructOf((*dns.DescribeDomainRecordsNewArgs)(nil)),
	"DescribeDomainRecordsNewResponse":    qlang.StructOf((*dns.DescribeDomainRecordsNewResponse)(nil)),
	"DescribeDomainRecordsResponse":       qlang.StructOf((*dns.DescribeDomainRecordsResponse)(nil)),
	"DescribeDomainsArgs":                 qlang.StructOf((*dns.DescribeDomainsArgs)(nil)),
	"DescribeSubDomainRecordsArgs":        qlang.StructOf((*dns.DescribeSubDomainRecordsArgs)(nil)),
	"DescribeSubDomainRecordsResponse":    qlang.StructOf((*dns.DescribeSubDomainRecordsResponse)(nil)),
	"DomainGroupType":                     qlang.StructOf((*dns.DomainGroupType)(nil)),
	"DomainType":                          qlang.StructOf((*dns.DomainType)(nil)),
	"GetMainDomainNameArgs":               qlang.StructOf((*dns.GetMainDomainNameArgs)(nil)),
	"GetMainDomainNameResponse":           qlang.StructOf((*dns.GetMainDomainNameResponse)(nil)),
	"RecordType":                          qlang.StructOf((*dns.RecordType)(nil)),
	"RecordTypeNew":                       qlang.StructOf((*dns.RecordTypeNew)(nil)),
	"UpdateDomainGroupArgs":               qlang.StructOf((*dns.UpdateDomainGroupArgs)(nil)),
	"UpdateDomainGroupResponse":           qlang.StructOf((*dns.UpdateDomainGroupResponse)(nil)),
	"UpdateDomainRecordArgs":              qlang.StructOf((*dns.UpdateDomainRecordArgs)(nil)),
	"UpdateDomainRecordResponse":          qlang.StructOf((*dns.UpdateDomainRecordResponse)(nil)),
}
