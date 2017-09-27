package crm

import (
	"github.com/denverdino/aliyungo/crm"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/denverdino/aliyungo/crm",

	"CRMAPIVersion":      crm.CRMAPIVersion,
	"CRMDefaultEndpoint": crm.CRMDefaultEndpoint,
	"FINANCE_LABEL":      crm.FINANCE_LABEL,
	"FINANCE_SERIES":     crm.FINANCE_SERIES,

	"Client":              qlang.StructOf((*crm.Client)(nil)),
	"client":              crm.NewClient,
	"CustomerLabel":       qlang.StructOf((*crm.CustomerLabel)(nil)),
	"LabelSeries":         qlang.StructOf((*crm.LabelSeries)(nil)),
	"LabelSeriesArgs":     qlang.StructOf((*crm.LabelSeriesArgs)(nil)),
	"LabelSeriesResponse": qlang.StructOf((*crm.LabelSeriesResponse)(nil)),
}
