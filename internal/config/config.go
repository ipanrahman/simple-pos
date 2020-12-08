package config

import "sync"

type Config struct {
	AppName    string
	AppVersion string
	AppHost    string
	AppPort    int
	Env        string
	Database   Database
}

var configInstance *Config
var once sync.Once

func NewConfig() *Config {
	once.Do(func() {
		configInstance = loadConfig()
	})
	return configInstance
}

func loadConfig() *Config {
	v := Viper()

	return &Config{
		AppName:    v.GetString("app.name"),
		AppVersion: v.GetString("app.version"),
		AppHost:    v.GetString("app.host"),
		AppPort:    v.GetInt("app.port"),
		Env:        v.GetString("app.env"),
		Database: Database{
			Driver:            v.GetString("database.postgres.driver"),
			Host:              v.GetString("database.postgres.host"),
			Port:              v.GetInt("database.postgres.port"),
			Username:          v.GetString("database.postgres.username"),
			Password:          v.GetString("database.postgres.password"),
			DatabaseName:      v.GetString("database.postgres.database_name"),
			MaxIdleConnection: v.GetInt("database.postgres.max_idle_connection"),
			MaxOpenConnection: v.GetInt("database.postgres.max_open_connection"),
		},
	}
}
