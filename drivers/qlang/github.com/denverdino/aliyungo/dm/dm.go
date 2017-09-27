package dm

import (
	"github.com/denverdino/aliyungo/dm"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/denverdino/aliyungo/dm",

	"BatchSendMail":   dm.BatchSendMail,
	"EmailAPIVersion": dm.EmailAPIVersion,
	"EmailEndPoint":   dm.EmailEndPoint,
	"SingleSendMail":  dm.SingleSendMail,

	"Client":             qlang.StructOf((*dm.Client)(nil)),
	"client":             dm.NewClient,
	"SendBatchMailArgs":  qlang.StructOf((*dm.SendBatchMailArgs)(nil)),
	"SendEmailArgs":      qlang.StructOf((*dm.SendEmailArgs)(nil)),
	"SendSingleMailArgs": qlang.StructOf((*dm.SendSingleMailArgs)(nil)),
}
