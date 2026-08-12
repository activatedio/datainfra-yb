package testing

import (
	"github.com/activatedio/datainfra-yb/pkg/data/yb"
	gormtesting "github.com/activatedio/datainfra/pkg/data/gorm/testing"
	datatesting "github.com/activatedio/datainfra/pkg/data/testing"
	gormmigrate "github.com/activatedio/datainfra/pkg/migrate/gorm"
	"go.uber.org/fx"
)

// NewAppFixture creates a new AppFixture for testing against YugabyteDB.
func NewAppFixture(name string, opt fx.Option) datatesting.AppFixture {
	return gormtesting.NewAppFixture(name, opt)
}

// NewStaticTestingConfig creates a static testing configuration function from YugabyteDB configs.
func NewStaticTestingConfig(ownerConfig, appConfig *yb.Config, migratorData []gormmigrate.MigratorData) func() gormtesting.GormTestingConfigResult {
	return gormtesting.NewStaticGormTestingConfig(yb.NewGormConfig(ownerConfig), yb.NewGormConfig(appConfig), migratorData)
}
