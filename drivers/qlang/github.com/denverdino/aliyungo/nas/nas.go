package nas

import (
	"github.com/denverdino/aliyungo/nas"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/denverdino/aliyungo/nas",

	"DEFAULT_POLICY":     nas.DEFAULT_POLICY,
	"DEFAULT_PRIORITY":   nas.DEFAULT_PRIORITY,
	"DEFAULT_SQUASHTYPE": nas.DEFAULT_SQUASHTYPE,
	"END_POINT":          nas.END_POINT,
	"VERSION":            nas.VERSION,

	"Client":                       qlang.StructOf((*nas.Client)(nil)),
	"client":                       nas.NewClient,
	"CreateAccessRuleRequest":      qlang.StructOf((*nas.CreateAccessRuleRequest)(nil)),
	"CreateAccessRuleResponse":     qlang.StructOf((*nas.CreateAccessRuleResponse)(nil)),
	"CreateFileSystemRequest":      qlang.StructOf((*nas.CreateFileSystemRequest)(nil)),
	"CreateFileSystemResponse":     qlang.StructOf((*nas.CreateFileSystemResponse)(nil)),
	"CreateMountTargetRequest":     qlang.StructOf((*nas.CreateMountTargetRequest)(nil)),
	"CreateMountTargetResponse":    qlang.StructOf((*nas.CreateMountTargetResponse)(nil)),
	"DescribeAccessRulesRequest":   qlang.StructOf((*nas.DescribeAccessRulesRequest)(nil)),
	"DescribeAccessRulesResponse":  qlang.StructOf((*nas.DescribeAccessRulesResponse)(nil)),
	"DescribeFileSystemsRequest":   qlang.StructOf((*nas.DescribeFileSystemsRequest)(nil)),
	"DescribeFileSystemsResponse":  qlang.StructOf((*nas.DescribeFileSystemsResponse)(nil)),
	"DescribeMountTargetsRequest":  qlang.StructOf((*nas.DescribeMountTargetsRequest)(nil)),
	"DescribeMountTargetsResponse": qlang.StructOf((*nas.DescribeMountTargetsResponse)(nil)),
	"FileSystem":                   qlang.StructOf((*nas.FileSystem)(nil)),
	"MountTarget":                  qlang.StructOf((*nas.MountTarget)(nil)),
	"Rule":                         qlang.StructOf((*nas.Rule)(nil)),
}
