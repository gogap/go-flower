package sts

import (
	"github.com/denverdino/aliyungo/sts"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/denverdino/aliyungo/sts",

	"STSAPIVersion":      sts.STSAPIVersion,
	"STSDefaultEndpoint": sts.STSDefaultEndpoint,

	"AssumeRoleRequest":          qlang.StructOf((*sts.AssumeRoleRequest)(nil)),
	"AssumeRoleResponse":         qlang.StructOf((*sts.AssumeRoleResponse)(nil)),
	"AssumedRoleUser":            qlang.StructOf((*sts.AssumedRoleUser)(nil)),
	"AssumedRoleUserCredentials": qlang.StructOf((*sts.AssumedRoleUserCredentials)(nil)),
	"STSClient":                  qlang.StructOf((*sts.STSClient)(nil)),
	"newClient":                  sts.NewClient,
	"newClientWithEndpoint":      sts.NewClientWithEndpoint,
}
