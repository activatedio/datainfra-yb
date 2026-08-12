package yb

import (
	"github.com/dave/jennifer/jen"
)

// ImportThis defines the import path for the yb runtime package used by generated code.
var (
	ImportThis = "github.com/activatedio/datainfra-yb/pkg/data/yb"
)

// TrigramSimilarityBinderCall returns jen code that evaluates to
// yb.TrigramSimilarityBinder(column, threshold), for use as a
// gorm.SearchBinding Binder.
func TrigramSimilarityBinderCall(column string, threshold float64) jen.Code {
	return jen.Qual(ImportThis, "TrigramSimilarityBinder").Call(jen.Lit(column), jen.Lit(threshold))
}
