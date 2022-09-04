package config

import (
	"MovieAPI/pkg/log"
	"github.com/spf13/viper"
	"sync"
)

var (
	singletonConfig sync.Once
	y               *configReader
)

// Config represents an application configuration.
type (
	Configuration struct {
		MongoSettings MongoSettings
	}
	MongoSettings struct {
		DatabaseName string
		Uri          string
	}
	Reader interface {
		GetAllValues() *Configuration
	}
	configReader struct {
		configFile string
		v          *viper.Viper
	}
)

func (y *configReader) GetAllValues() *Configuration {

	var configuration Configuration
	if err := y.v.ReadInConfig(); err != nil {
		log.Logger.Error(err)
		panic(err)
	}
	err := y.v.Unmarshal(&configuration)
	if err != nil {
		log.Logger.Error(err)
		panic(err)
	}
	return &configuration
}

func NewConfigReader(configPath string, configFile string, vip *viper.Viper) Reader {
	singletonConfig.Do(func() {
		vip.SetConfigName(configFile)
		vip.AddConfigPath(configPath)
		y = &configReader{
			configFile: configFile,
			v:          vip,
		}
	})
	return y
}
