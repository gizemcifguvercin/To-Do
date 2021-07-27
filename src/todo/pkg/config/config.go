package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var cfgReader *configReader

const (
	defaultEnv = "dev"
)

type (
	configReader struct {
		configFile string
		v          *viper.Viper
	}
)

func GetValueByKey(key string) string {
	newConfigReader()

	var err error
	if err = cfgReader.v.ReadInConfig(); err != nil {
		fmt.Printf("Failed to read config file : %s", err)
		return  err.Error()
	}
	return cfgReader.v.GetString(key)
}

func getEnvironment(env, fallback string) string{
	e := os.Getenv(env)
	if e == ""{
		return fallback
	}
	return e
}

func newConfigReader() {
	env := getEnvironment("APP_ENVIRONMENT", defaultEnv)
	configFile := fmt.Sprintf("api.%s.yaml",env)
	v := viper.GetViper()
	v.SetConfigType("yaml")
	v.SetConfigFile(configFile)
	cfgReader = &configReader{
		configFile: configFile,
		v:          v,
	}
}

