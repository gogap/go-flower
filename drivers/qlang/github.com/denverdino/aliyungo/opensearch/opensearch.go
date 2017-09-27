package opensearch

import (
	"github.com/denverdino/aliyungo/opensearch"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/denverdino/aliyungo/opensearch",

	"APIVersion": opensearch.APIVersion,
	"Internet":   opensearch.Internet,
	"Intranet":   opensearch.Intranet,
	"VPC":        opensearch.VPC,

	"Client":         qlang.StructOf((*opensearch.Client)(nil)),
	"client":         opensearch.NewClient,
	"OpenSearchArgs": qlang.StructOf((*opensearch.OpenSearchArgs)(nil)),
	"PushArgs":       qlang.StructOf((*opensearch.PushArgs)(nil)),
	"SearchArgs":     qlang.StructOf((*opensearch.SearchArgs)(nil)),
}
