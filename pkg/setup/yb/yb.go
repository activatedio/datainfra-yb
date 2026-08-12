package yb

import (
	"github.com/activatedio/datainfra-yb/pkg/data/yb"
	gormsetup "github.com/activatedio/datainfra/pkg/setup/gorm"
)

// OwnerConfig has the configuration for the database owner
type OwnerConfig struct {
	Config yb.Config
}

// NewOwnerGormConfig maps a YugabyteDB owner configuration to the gorm setup configuration.
func NewOwnerGormConfig(config *OwnerConfig) *gormsetup.OwnerGormConfig {
	return &gormsetup.OwnerGormConfig{
		Config: *yb.NewGormConfig(&config.Config),
	}
}
