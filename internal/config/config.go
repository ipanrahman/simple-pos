package config

type Config struct {
	AppName    string
	AppVersion string
	AppHost    string
	AppPort    int
	Env        string
	Database   Database
}

func NewConfig() *Config {
	v := Viper()
	return &Config{
		AppName:    v.GetString("app.name"),
		AppVersion: v.GetString("app.version"),
		AppHost:    v.GetString("app.host"),
		AppPort:    v.GetInt("app.port"),
		Env:        v.GetString("app.env"),
		Database: Database{
			Driver:            v.GetString("db.driver"),
			Host:              v.GetString("db.host"),
			Port:              v.GetInt("db.port"),
			Username:          v.GetString("db.user"),
			Password:          v.GetString("db.password"),
			DatabaseName:      v.GetString("db.database"),
			MaxIdleConnection: v.GetInt("db.max_idle_connection"),
			MaxOpenConnection: v.GetInt("db.max_open_connection"),
		},
	}
}
