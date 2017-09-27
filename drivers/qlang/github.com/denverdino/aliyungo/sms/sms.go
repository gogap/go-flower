package sms

import (
	"github.com/denverdino/aliyungo/sms"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/denverdino/aliyungo/sms",

	"DYSmsAPIVersion": sms.DYSmsAPIVersion,
	"DYSmsEndPoint":   sms.DYSmsEndPoint,
	"SendSms":         sms.SendSms,
	"SingleSendSms":   sms.SingleSendSms,
	"SmsAPIVersion":   sms.SmsAPIVersion,
	"SmsEndPoint":     sms.SmsEndPoint,

	"Client":            qlang.StructOf((*sms.Client)(nil)),
	"client":            sms.NewClient,
	"DYSmsClient":       qlang.StructOf((*sms.DYSmsClient)(nil)),
	"dysmsclient":       sms.NewDYSmsClient,
	"SendSmsArgs":       qlang.StructOf((*sms.SendSmsArgs)(nil)),
	"SendSmsResponse":   qlang.StructOf((*sms.SendSmsResponse)(nil)),
	"SingleSendSmsArgs": qlang.StructOf((*sms.SingleSendSmsArgs)(nil)),
}
