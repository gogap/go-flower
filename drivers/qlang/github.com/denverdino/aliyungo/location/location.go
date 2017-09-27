package location

import (
	"github.com/denverdino/aliyungo/location"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/denverdino/aliyungo/location",

	"LocationAPIVersion":      location.LocationAPIVersion,
	"LocationDefaultEndpoint": location.LocationDefaultEndpoint,

	"Client":                    qlang.StructOf((*location.Client)(nil)),
	"newClient":                 location.NewClient,
	"newClientWithEndpoint":     location.NewClientWithEndpoint,
	"DescribeEndpointArgs":      qlang.StructOf((*location.DescribeEndpointArgs)(nil)),
	"DescribeEndpointResponse":  qlang.StructOf((*location.DescribeEndpointResponse)(nil)),
	"DescribeEndpointsArgs":     qlang.StructOf((*location.DescribeEndpointsArgs)(nil)),
	"DescribeEndpointsResponse": qlang.StructOf((*location.DescribeEndpointsResponse)(nil)),
	"DescribeRegionsArgs":       qlang.StructOf((*location.DescribeRegionsArgs)(nil)),
	"DescribeRegionsResponse":   qlang.StructOf((*location.DescribeRegionsResponse)(nil)),
	"DescribeServicesArgs":      qlang.StructOf((*location.DescribeServicesArgs)(nil)),
	"DescribeServicesResponse":  qlang.StructOf((*location.DescribeServicesResponse)(nil)),
	"EndpointItem":              qlang.StructOf((*location.EndpointItem)(nil)),
}
