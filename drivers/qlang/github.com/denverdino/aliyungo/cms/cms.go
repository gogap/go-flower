package cms

import (
	"github.com/denverdino/aliyungo/cms"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/denverdino/aliyungo/cms",

	"ACS_PREFIX":         cms.ACS_PREFIX,
	"APIVersion":         cms.APIVersion,
	"CMSAPIVersion":      cms.CMSAPIVersion,
	"CMSDefaultEndpoint": cms.CMSDefaultEndpoint,
	"CMSServiceCode":     cms.CMSServiceCode,
	"DefaultEndpoint":    cms.DefaultEndpoint,
	"HEADER_SEPERATER":   cms.HEADER_SEPERATER,
	"METHOD_DELETE":      cms.METHOD_DELETE,
	"METHOD_GET":         cms.METHOD_GET,
	"METHOD_POST":        cms.METHOD_POST,
	"METHOD_PUT":         cms.METHOD_PUT,

	"bodyMd5":        cms.BodyMd5,
	"getRequestPath": cms.GetRequestPath,
	"initBaseHeader": cms.InitBaseHeader,

	"ActionGroup":              qlang.StructOf((*cms.ActionGroup)(nil)),
	"ActionsModel":             qlang.StructOf((*cms.ActionsModel)(nil)),
	"AlarmHistoryItem":         qlang.StructOf((*cms.AlarmHistoryItem)(nil)),
	"AlarmItem":                qlang.StructOf((*cms.AlarmItem)(nil)),
	"AlertActionsModel":        qlang.StructOf((*cms.AlertActionsModel)(nil)),
	"AlertRequest":             qlang.StructOf((*cms.AlertRequest)(nil)),
	"CMSClient":                qlang.StructOf((*cms.CMSClient)(nil)),
	"cmsclient":                cms.NewCMSClient,
	"newCMSRegionClient":       cms.NewCMSRegionClient,
	"newClientWithEndpoint":    cms.NewClientWithEndpoint,
	"newClientWithRegion":      cms.NewClientWithRegion,
	"client":                   cms.NewClient,
	"CommonAlarmItem":          qlang.StructOf((*cms.CommonAlarmItem)(nil)),
	"CommonAlarmResponse":      qlang.StructOf((*cms.CommonAlarmResponse)(nil)),
	"ConditionModel":           qlang.StructOf((*cms.ConditionModel)(nil)),
	"CreateAlarmArgs":          qlang.StructOf((*cms.CreateAlarmArgs)(nil)),
	"CreateAlarmResponse":      qlang.StructOf((*cms.CreateAlarmResponse)(nil)),
	"DeleteAlarmArgs":          qlang.StructOf((*cms.DeleteAlarmArgs)(nil)),
	"DeleteAlarmResponse":      qlang.StructOf((*cms.DeleteAlarmResponse)(nil)),
	"DimensionDataPoint":       qlang.StructOf((*cms.DimensionDataPoint)(nil)),
	"DimensionRequest":         qlang.StructOf((*cms.DimensionRequest)(nil)),
	"DisableAlarmArgs":         qlang.StructOf((*cms.DisableAlarmArgs)(nil)),
	"DisableAlarmResponse":     qlang.StructOf((*cms.DisableAlarmResponse)(nil)),
	"EnableAlarmArgss":         qlang.StructOf((*cms.EnableAlarmArgss)(nil)),
	"EnableAlarmResponse":      qlang.StructOf((*cms.EnableAlarmResponse)(nil)),
	"EscalationsModel":         qlang.StructOf((*cms.EscalationsModel)(nil)),
	"Failure":                  qlang.StructOf((*cms.Failure)(nil)),
	"GetDimenstionResult":      qlang.StructOf((*cms.GetDimenstionResult)(nil)),
	"GetProjectResult":         qlang.StructOf((*cms.GetProjectResult)(nil)),
	"ListAlarmArgs":            qlang.StructOf((*cms.ListAlarmArgs)(nil)),
	"ListAlarmHistoryArgs":     qlang.StructOf((*cms.ListAlarmHistoryArgs)(nil)),
	"ListAlarmHistoryResponse": qlang.StructOf((*cms.ListAlarmHistoryResponse)(nil)),
	"ListAlarmResponse":        qlang.StructOf((*cms.ListAlarmResponse)(nil)),
	"ListContactGroupArgs":     qlang.StructOf((*cms.ListContactGroupArgs)(nil)),
	"ListContactGroupResponse": qlang.StructOf((*cms.ListContactGroupResponse)(nil)),
	"ListProjectRequestModel":  qlang.StructOf((*cms.ListProjectRequestModel)(nil)),
	"ProjectListResultModel":   qlang.StructOf((*cms.ProjectListResultModel)(nil)),
	"ProjectModel":             qlang.StructOf((*cms.ProjectModel)(nil)),
	"ProjectResultModel":       qlang.StructOf((*cms.ProjectResultModel)(nil)),
	"ResultModel":              qlang.StructOf((*cms.ResultModel)(nil)),
	"UpdateAlarmArgs":          qlang.StructOf((*cms.UpdateAlarmArgs)(nil)),
	"UpdateAlarmResponse":      qlang.StructOf((*cms.UpdateAlarmResponse)(nil)),
}
