package config

import "time"

const (
	DefaultMaxIdleConnection     = 5
	DefaultMaxOpenConnection     = 5
	DefaultConnectionMaxLifeTime = 10 * time.Minute
)
