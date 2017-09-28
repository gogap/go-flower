package cs

import (
	"github.com/denverdino/aliyungo/cs"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/denverdino/aliyungo/cs",

	"CSAPIVersion":           cs.CSAPIVersion,
	"CSDefaultEndpoint":      cs.CSDefaultEndpoint,
	"ClassicNetwork":         cs.ClassicNetwork,
	"ClusterDefaultTimeout":  cs.ClusterDefaultTimeout,
	"DefaultPreSleepTime":    cs.DefaultPreSleepTime,
	"DefaultWaitForInterval": cs.DefaultWaitForInterval,
	"DeleteFailed":           cs.DeleteFailed,
	"Deleted":                cs.Deleted,
	"Deleting":               cs.Deleting,
	"Failed":                 cs.Failed,
	"InActive":               cs.InActive,
	"Initial":                cs.Initial,
	"NASDriver":              cs.NASDriver,
	"OSSFSDriver":            cs.OSSFSDriver,
	"Running":                cs.Running,
	"ScaleTo":                cs.ScaleTo,
	"Scaling":                cs.Scaling,
	"Updating":               cs.Updating,
	"VPCNetwork":             cs.VPCNetwork,

	"client":                  cs.NewClient,
	"ClusterCerts":            qlang.StructOf((*cs.ClusterCerts)(nil)),
	"ClusterCreationArgs":     qlang.StructOf((*cs.ClusterCreationArgs)(nil)),
	"ClusterCreationResponse": qlang.StructOf((*cs.ClusterCreationResponse)(nil)),
	"ClusterResizeArgs":       qlang.StructOf((*cs.ClusterResizeArgs)(nil)),
	"ClusterType":             qlang.StructOf((*cs.ClusterType)(nil)),
	"Container":               qlang.StructOf((*cs.Container)(nil)),
	"Definition":              qlang.StructOf((*cs.Definition)(nil)),
	"GetVolumeResponse":       qlang.StructOf((*cs.GetVolumeResponse)(nil)),
	"GetVolumesResponse":      qlang.StructOf((*cs.GetVolumesResponse)(nil)),
	"NASOpts":                 qlang.StructOf((*cs.NASOpts)(nil)),
	"NodeStatus":              qlang.StructOf((*cs.NodeStatus)(nil)),
	"OSSOpts":                 qlang.StructOf((*cs.OSSOpts)(nil)),
	"Port":                    qlang.StructOf((*cs.Port)(nil)),
	"Project":                 qlang.StructOf((*cs.Project)(nil)),
	"projectClient":           cs.NewProjectClient,
	"ProjectCreationArgs":     qlang.StructOf((*cs.ProjectCreationArgs)(nil)),
	"ProjectUpdationArgs":     qlang.StructOf((*cs.ProjectUpdationArgs)(nil)),
	"Request":                 qlang.StructOf((*cs.Request)(nil)),
	"Response":                qlang.StructOf((*cs.Response)(nil)),
	"ScaleServiceArgs":        qlang.StructOf((*cs.ScaleServiceArgs)(nil)),
	"Service":                 qlang.StructOf((*cs.Service)(nil)),
	"VolumeCreationArgs":      qlang.StructOf((*cs.VolumeCreationArgs)(nil)),
	"VolumeCreationResponse":  qlang.StructOf((*cs.VolumeCreationResponse)(nil)),
	"VolumeRef":               qlang.StructOf((*cs.VolumeRef)(nil)),

	"NetworkModeType":  qlang.StructOf((*cs.NetworkModeType)(nil)),
	"ClusterState":     qlang.StructOf((*cs.ClusterState)(nil)),
	"VolumeDriverType": qlang.StructOf((*cs.VolumeDriverType)(nil)),
}
