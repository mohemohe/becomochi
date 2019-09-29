package configs

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	environmentVariable struct {
		Echo struct{
			Env string `default:"debug"`
		}
		Mongo struct {
			Address  string `default:"mongodb://localhost:27017"`
			Database string `default:"parakeet"`
			SSL      bool   `default:"false"`
		}
		Becomochi struct {
			Api struct {
				Key string `default:"root"`
			}
		}
	}
)

var cachedEnvironmentVariable *environmentVariable

func GetEnv() environmentVariable {
	if cachedEnvironmentVariable == nil {
		cachedEnvironmentVariable = new(environmentVariable)
		envconfig.MustProcess("", cachedEnvironmentVariable)
	}

	return *cachedEnvironmentVariable
}
