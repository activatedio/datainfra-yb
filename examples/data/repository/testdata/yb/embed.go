// Package yb contains YugabyteDB related test data
package yb

import "embed"

// Files contains test data migration files
//
//go:embed *.sql
var Files embed.FS
