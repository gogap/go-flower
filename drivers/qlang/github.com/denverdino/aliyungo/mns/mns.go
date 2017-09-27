package mns

import (
	"github.com/denverdino/aliyungo/mns"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/denverdino/aliyungo/mns",

	"HeaderMNSPrefix": mns.HeaderMNSPrefix,
	"MNSAPIVersion":   mns.MNSAPIVersion,

	"getCurrentMillisecond": mns.GetCurrentMillisecond,
	"getCurrentUnixMicro":   mns.GetCurrentUnixMicro,
	"hamSha1":               mns.HamSha1,
	"md5":                   mns.Md5,
	"sha1":                  mns.Sha1,

	"client":     mns.NewClient,
	"Message":    qlang.StructOf((*mns.Message)(nil)),
	"MsgReceive": qlang.StructOf((*mns.MsgReceive)(nil)),
	"MsgSend":    qlang.StructOf((*mns.MsgSend)(nil)),
	"Queue":      qlang.StructOf((*mns.Queue)(nil)),
}
