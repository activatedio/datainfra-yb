package yb

import (
	"strings"

	datagorm "github.com/activatedio/datainfra/pkg/data/gorm"
	"gorm.io/gorm"
)

// NewGormConfig maps a YugabyteDB Config to the gorm data stack configuration.
// YSQL is PostgreSQL-compatible, so the postgres dialect is used. Multiple
// hosts are passed as a comma-separated host list, which the underlying pgx
// driver resolves with client-side failover.
func NewGormConfig(config *Config) *datagorm.Config {

	port := config.Port
	if port == 0 {
		port = DefaultPort
	}

	return &datagorm.Config{
		Dialect:                  datagorm.DialectPostgres,
		Host:                     strings.Join(config.Hosts, ","),
		Port:                     port,
		Username:                 config.Username,
		Password:                 config.Password,
		Name:                     config.Name,
		EnableDefaultTransaction: config.EnableDefaultTransaction,
		EnableSQLLogging:         config.EnableSQLLogging,
		MaxIdleConns:             config.MaxIdleConns,
	}
}

// NewDB creates a new gorm database instance for the given YugabyteDB configuration.
func NewDB(config *Config) (*gorm.DB, error) {
	return datagorm.NewDB(NewGormConfig(config))
}
