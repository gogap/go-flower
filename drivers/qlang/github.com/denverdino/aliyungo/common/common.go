package common

import (
	"github.com/denverdino/aliyungo/common"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/denverdino/aliyungo/common",

	"APNorthEast1":       common.APNorthEast1,
	"APSouthEast1":       common.APSouthEast1,
	"APSouthEast2":       common.APSouthEast2,
	"Beijing":            common.Beijing,
	"Classic":            common.Classic,
	"Day":                common.Day,
	"ECSRequestMethod":   common.ECSRequestMethod,
	"EUCentral1":         common.EUCentral1,
	"HTTPS_PROTOCOL":     common.HTTPS_PROTOCOL,
	"HTTP_PROTOCOL":      common.HTTP_PROTOCOL,
	"Hangzhou":           common.Hangzhou,
	"Hongkong":           common.Hongkong,
	"Hour":               common.Hour,
	"Internet":           common.Internet,
	"Intranet":           common.Intranet,
	"JSONResponseFormat": common.JSONResponseFormat,
	"MEEast1":            common.MEEast1,
	"Month":              common.Month,
	"PayByBandwidth":     common.PayByBandwidth,
	"PayByTraffic":       common.PayByTraffic,
	"PostPaid":           common.PostPaid,
	"PrePaid":            common.PrePaid,
	"Qingdao":            common.Qingdao,
	"Shanghai":           common.Shanghai,
	"Shenzhen":           common.Shenzhen,
	"SignatureMethod":    common.SignatureMethod,
	"SignatureVersion":   common.SignatureVersion,
	"USEast1":            common.USEast1,
	"USWest1":            common.USWest1,
	"VPC":                common.VPC,
	"Version":            common.Version,
	"XMLResponseFormat":  common.XMLResponseFormat,
	"Year":               common.Year,
	"Zhangjiakou":        common.Zhangjiakou,
	"Region":             func(regionId string) common.Region { return common.Region(regionId) },

	"ValidRegions": common.ValidRegions,

	"getClientError":           common.GetClientError,
	"getClientErrorFromString": common.GetClientErrorFromString,

	"BusinessInfo":             qlang.StructOf((*common.BusinessInfo)(nil)),
	"newLocationClient":        common.NewLocationClient,
	"DescribeEndpointArgs":     qlang.StructOf((*common.DescribeEndpointArgs)(nil)),
	"DescribeEndpointResponse": qlang.StructOf((*common.DescribeEndpointResponse)(nil)),
	"Endpoint":                 qlang.StructOf((*common.Endpoint)(nil)),
	"EndpointItem":             qlang.StructOf((*common.EndpointItem)(nil)),
	"Endpoints":                qlang.StructOf((*common.Endpoints)(nil)),
	"ErrorResponse":            qlang.StructOf((*common.ErrorResponse)(nil)),
	"Pagination":               qlang.StructOf((*common.Pagination)(nil)),
	"PaginationResult":         qlang.StructOf((*common.PaginationResult)(nil)),
	"Product":                  qlang.StructOf((*common.Product)(nil)),
	"Products":                 qlang.StructOf((*common.Products)(nil)),
	"RegionIds":                qlang.StructOf((*common.RegionIds)(nil)),
	"Request":                  qlang.StructOf((*common.Request)(nil)),
	"Response":                 qlang.StructOf((*common.Response)(nil)),
}
