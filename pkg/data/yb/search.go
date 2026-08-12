package yb

import (
	"fmt"

	"github.com/activatedio/datainfra/pkg/data"
	datagorm "github.com/activatedio/datainfra/pkg/data/gorm"
	"gorm.io/gorm"
)

// TrigramSimilarityBinder returns a PredicateBinder that filters using
// pg_trgm trigram similarity on the given column, keeping rows whose
// similarity to the predicate StringValue exceeds the threshold. This gives
// typo-tolerant matching entirely inside the database. The predicate
// operator must be SearchOperatorStringMatch.
//
// Requires the pg_trgm extension. A `USING ybgin (<column> gin_trgm_ops)`
// index accelerates the lookup on YugabyteDB.
func TrigramSimilarityBinder(column string, threshold float64) datagorm.PredicateBinder {
	return func(tx *gorm.DB, p *data.SearchPredicate) (*gorm.DB, error) {
		if p.Operator != data.SearchOperatorStringMatch {
			return nil, fmt.Errorf("TrigramSimilarityBinder: operator %v not supported", p.Operator)
		}
		return tx.Where(fmt.Sprintf("similarity(%s, ?) > ?", column), p.StringValue, threshold), nil
	}
}
