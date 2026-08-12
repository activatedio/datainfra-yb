package yb

import (
	gormmigrate "github.com/activatedio/datainfra/pkg/migrate/gorm"

	"github.com/activatedio/datainfra-yb/pkg/data/yb"
)

// MigratorConfig defines the configuration for the YugabyteDB migrator.
type MigratorConfig struct {
	Config yb.Config
}

// NewMigratorGormConfig maps a YugabyteDB migrator configuration to the gorm migrator configuration.
func NewMigratorGormConfig(config *MigratorConfig) *gormmigrate.MigratorGormConfig {
	return &gormmigrate.MigratorGormConfig{
		GormConfig: *yb.NewGormConfig(&config.Config),
	}
}
