package orm

import (
	"testing"

	"github.com/sherlockjjobs/sigma/framework/contract"
	"github.com/sherlockjjobs/sigma/framework/provider/config"
	tests "github.com/sherlockjjobs/sigma/test"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHadeConfig_Load(t *testing.T) {
	container := tests.InitBaseContainer()
	container.Bind(&config.HadeConfigProvider{})

	Convey("test config", t, func() {
		configService := container.MustMake(contract.ConfigKey).(contract.Config)
		config := &contract.DBConfig{}
		err := configService.Load("database.default", config)
		So(err, ShouldBeNil)
	})

	Convey("test default config", t, func() {
		configService := container.MustMake(contract.ConfigKey).(contract.Config)
		config := &contract.DBConfig{
			ConnMaxIdle: 10,
		}
		err := configService.Load("database.read", config)
		So(err, ShouldBeNil)
		So(config.ConnMaxIdle, ShouldEqual, 10)
	})

	Convey("test base config", t, func() {
		configService := container.MustMake(contract.ConfigKey).(contract.Config)
		config := &contract.DBConfig{
			ConnMaxOpen: 200,
		}
		err := configService.Load("database", config)
		So(err, ShouldBeNil)
		So(config.ConnMaxOpen, ShouldEqual, 100)
	})

}
