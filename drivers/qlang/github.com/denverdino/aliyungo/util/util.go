package util

import (
	"github.com/denverdino/aliyungo/util"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/denverdino/aliyungo/util",

	"convertToQueryValues":         util.ConvertToQueryValues,
	"createRandomString":           util.CreateRandomString,
	"createSignature":              util.CreateSignature,
	"createSignatureForRequest":    util.CreateSignatureForRequest,
	"encode":                       util.Encode,
	"flattenFn":                    util.FlattenFn,
	"generateRandomECSPassword":    util.GenerateRandomECSPassword,
	"getGMTime":                    util.GetGMTime,
	"getISO8601TimeStamp":          util.GetISO8601TimeStamp,
	"setQueryValueByFlattenMethod": util.SetQueryValueByFlattenMethod,
	"setQueryValues":               util.SetQueryValues,
	"underline2Dot":                util.Underline2Dot,

	"AttemptStrategy": qlang.StructOf((*util.AttemptStrategy)(nil)),
}
