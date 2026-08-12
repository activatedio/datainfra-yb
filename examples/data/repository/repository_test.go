package repository_test

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	ybtestdata "github.com/activatedio/datainfra-yb/examples/data/repository/testdata/yb"
	repoyb "github.com/activatedio/datainfra-yb/examples/data/repository/yb"
	ybmigrations "github.com/activatedio/datainfra-yb/examples/data/repository/yb/migrations"
	"github.com/activatedio/datainfra-yb/pkg/data/yb"
	ybtesting "github.com/activatedio/datainfra-yb/pkg/data/yb/testing"
	datatesting "github.com/activatedio/datainfra/pkg/data/testing"
	gormmigrate "github.com/activatedio/datainfra/pkg/migrate/gorm"
	"go.uber.org/fx"
)

var (
	AppFixtures []datatesting.AppFixture
)

func TestMain(m *testing.M) {

	name := fmt.Sprintf("unit_%d_%d", time.Now().UnixNano(), os.Getpid())

	ybHosts := os.Getenv("YB_HOSTS")

	if ybHosts == "" {
		ybHosts = "127.0.0.1"
	}

	hosts := strings.Split(ybHosts, ",")

	migratorData := []gormmigrate.MigratorData{
		{
			Name: "main",
			FS:   ybmigrations.Files,
			Path: ".",
		},
		{
			Name: "test",
			FS:   ybtestdata.Files,
			Path: ".",
		},
	}

	zero := 0

	AppFixtures = []datatesting.AppFixture{
		ybtesting.NewAppFixture("yugabyte", fx.Module("testing", repoyb.Index(),
			fx.Provide(ybtesting.NewStaticTestingConfig(&yb.Config{
				Hosts:                    hosts,
				Username:                 "yugabyte",
				Password:                 "yugabyte",
				Name:                     "yugabyte",
				EnableDefaultTransaction: true,
				EnableSQLLogging:         true,
				MaxIdleConns:             &zero,
			}, &yb.Config{
				Hosts:                    hosts,
				Username:                 name,
				Password:                 name,
				Name:                     name,
				EnableDefaultTransaction: true,
				EnableSQLLogging:         true,
				MaxIdleConns:             &zero,
			}, migratorData)))),
	}

	rc := m.Run()

	for _, fixture := range AppFixtures {
		if err := fixture.Cleanup(); err != nil {
			panic(err)
		}
	}

	os.Exit(rc)

}
