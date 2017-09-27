package structs

import (
	"github.com/fatih/structs"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/fatih/structs",

	"DefaultTagName": structs.DefaultTagName,

	"fields":   structs.Fields,
	"fillMap":  structs.FillMap,
	"hasZero":  structs.HasZero,
	"isStruct": structs.IsStruct,
	"isZero":   structs.IsZero,
	"map":      structs.Map,
	"name":     structs.Name,
	"names":    structs.Names,
	"values":   structs.Values,

	"new": structs.New,
}
