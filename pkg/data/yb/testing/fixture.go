package testing

import (
	"github.com/activatedio/datainfra-yb/pkg/data/yb"
	gormtesting "github.com/activatedio/datainfra/pkg/data/gorm/testing"
	datatesting "github.com/activatedio/datainfra/pkg/data/testing"
	"go.uber.org/fx"
)

// NewAppFixture creates a new AppFixture for testing against YugabyteDB.
func NewAppFixture(name string, opt fx.Option) datatesting.LifecycleFixture {
	return gormtesting.NewAppFixture(name, opt)
}

// NewStaticTestingConfig creates a static testing configuration function from YugabyteDB configs.
func NewStaticTestingConfig(ownerConfig, appConfig *yb.Config) func() gormtesting.GormTestingConfigResult {
	return gormtesting.NewStaticGormTestingConfig(yb.NewGormConfig(ownerConfig), yb.NewGormConfig(appConfig))
}
