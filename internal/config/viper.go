package config

import (
	"github.com/spf13/viper"
	"log"
	"strings"
	"sync"
)

const (
	configFileName = "config"
	configFileExt  = "toml"
)

var viperInstance *viper.Viper
var viperInit sync.Once

// Viper load config
func Viper() *viper.Viper {
	viperInit.Do(func() {
		viperInstance = viper.New()
		viperInstance.AutomaticEnv()
		viperInstance.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viperInstance.SetConfigName(configFileName)
		viperInstance.AddConfigPath(".")
		viperInstance.SetConfigType(configFileExt)

		// Set Default
		viperInstance.SetDefault("app.port", 3000)
		// Read Config
		errReadConfig := viperInstance.ReadInConfig()
		if errReadConfig != nil {
			log.Fatalf("Error read config file : %v", errReadConfig)
		}
	})
	return viperInstance
}
