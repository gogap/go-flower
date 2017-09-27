package sls

import (
	"github.com/denverdino/aliyungo/sls"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/denverdino/aliyungo/sls",

	"HeaderSLSPrefix1":   sls.HeaderSLSPrefix1,
	"HeaderSLSPrefix2":   sls.HeaderSLSPrefix2,
	"METHOD_DELETE":      sls.METHOD_DELETE,
	"METHOD_GET":         sls.METHOD_GET,
	"METHOD_POST":        sls.METHOD_POST,
	"METHOD_PUT":         sls.METHOD_PUT,
	"SLSAPIVersion":      sls.SLSAPIVersion,
	"SLSDefaultEndpoint": sls.SLSDefaultEndpoint,

	"newClient":             sls.NewClient,
	"newClientWithEndpoint": sls.NewClientWithEndpoint,
	"GroupAttribute":        qlang.StructOf((*sls.GroupAttribute)(nil)),
	"IndexConfig":           qlang.StructOf((*sls.IndexConfig)(nil)),
	"IndexKeyConfig":        qlang.StructOf((*sls.IndexKeyConfig)(nil)),
	"IndexLineConfig":       qlang.StructOf((*sls.IndexLineConfig)(nil)),
	"Log":                   qlang.StructOf((*sls.Log)(nil)),
	"LogGroup":              qlang.StructOf((*sls.LogGroup)(nil)),
	"LogGroupList":          qlang.StructOf((*sls.LogGroupList)(nil)),
	"Log_Content":           qlang.StructOf((*sls.Log_Content)(nil)),
	"Logstore":              qlang.StructOf((*sls.Logstore)(nil)),
	"LogstoreList":          qlang.StructOf((*sls.LogstoreList)(nil)),
	"LogtailConfig":         qlang.StructOf((*sls.LogtailConfig)(nil)),
	"LogtailConfigList":     qlang.StructOf((*sls.LogtailConfigList)(nil)),
	"LogtailInput":          qlang.StructOf((*sls.LogtailInput)(nil)),
	"LogtailOutput":         qlang.StructOf((*sls.LogtailOutput)(nil)),
	"Machine":               qlang.StructOf((*sls.Machine)(nil)),
	"MachineGroup":          qlang.StructOf((*sls.MachineGroup)(nil)),
	"MachineGroupList":      qlang.StructOf((*sls.MachineGroupList)(nil)),
	"MachineList":           qlang.StructOf((*sls.MachineList)(nil)),
	"PutLogsRequest":        qlang.StructOf((*sls.PutLogsRequest)(nil)),
}
