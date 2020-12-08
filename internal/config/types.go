package config

import "time"

const (
	DefaultMaxIdleConnection     = 5
	DefaultMaxOpenConnection     = 5
	DefaultConnectionMaxLifeTime = 10 * time.Minute
)

type Database struct {
	Driver            string `json:"driver"`
	Host              string `json:"host"`
	Port              int    `json:"port"`
	Username          string `json:"username"`
	Password          string `json:"password"`
	DatabaseName      string `json:"database_name"`
	MaxIdleConnection int    `json:"max_idle_connection"`
	MaxOpenConnection int    `json:"max_open_connection"`
}
