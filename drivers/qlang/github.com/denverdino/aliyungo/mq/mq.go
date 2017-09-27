package mq

import (
	"github.com/denverdino/aliyungo/mq"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/denverdino/aliyungo/mq",

	"getCurrentMillisecond": mq.GetCurrentMillisecond,
	"getCurrentUnixMicro":   mq.GetCurrentUnixMicro,
	"hamSha1":               mq.HamSha1,
	"httpDelete":            mq.HttpDelete,
	"httpGet":               mq.HttpGet,
	"md5":                   mq.Md5,
	"sha1":                  mq.Sha1,

	"Client":          qlang.StructOf((*mq.Client)(nil)),
	"client":          mq.NewClient,
	"Message":         qlang.StructOf((*mq.Message)(nil)),
	"MessageResponse": qlang.StructOf((*mq.MessageResponse)(nil)),
	"Messages":        qlang.StructOf((*mq.Messages)(nil)),
	"SendMessage":     qlang.StructOf((*mq.SendMessage)(nil)),
}
