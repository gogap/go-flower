package pwgen

import (
	"github.com/chr4/pwgen"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/chr4/pwgen",

	"alphaNum":        pwgen.AlphaNum,
	"alphaNumSymbols": pwgen.AlphaNumSymbols,
	"new":             pwgen.New,
	"num":             pwgen.Num,
}
