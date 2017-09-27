package push

import (
	"github.com/denverdino/aliyungo/push"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/denverdino/aliyungo/push",

	"Push":                     push.Push,
	"PushAPIVersion":           push.PushAPIVersion,
	"PushDeviceTypeAll":        push.PushDeviceTypeAll,
	"PushDeviceTypeAndroid":    push.PushDeviceTypeAndroid,
	"PushDeviceTypeIOS":        push.PushDeviceTypeIOS,
	"PushEndPoint":             push.PushEndPoint,
	"PushIOSAPNENVDevelopment": push.PushIOSAPNENVDevelopment,
	"PushIOSAPNENVProduct":     push.PushIOSAPNENVProduct,
	"PushTargetAccount":        push.PushTargetAccount,
	"PushTargetAlias":          push.PushTargetAlias,
	"PushTargetAll":            push.PushTargetAll,
	"PushTargetDevice":         push.PushTargetDevice,
	"PushTargetTag":            push.PushTargetTag,
	"PushTypeMessage":          push.PushTypeMessage,
	"PushTypeNotice":           push.PushTypeNotice,

	"Client":       qlang.StructOf((*push.Client)(nil)),
	"client":       push.NewClient,
	"PushArgs":     qlang.StructOf((*push.PushArgs)(nil)),
	"PushResponse": qlang.StructOf((*push.PushResponse)(nil)),
}
