package db

import "time"

const (
	// MysqlConnectionDSNFormat is mysql connection path format for gorm.
	// E.g. app:password@tcp(localhost:3306)/app?charset=utf8&parseTime=True&loc=Local
	MysqlConnectionDSNFormat = "%s:%s@%s(%s)/%s?charset=utf8&parseTime=True&loc=Local"
)

// Config struct holds db configurations and implements ConfigReader.
type Config struct {
	Dialect               string `toml:"Dialect"`
	Protocol              string `toml:"Protocol"`
	URL                   string `toml:"URL"`
	Username              string `toml:"Username"`
	Password              string `env:"Password"`
	SslMode               string `toml:"SslMode"`
	Name                  string `toml:"Name"`
	AltersEnabled         bool   `toml:"AltersEnabled"`
	MaxOpenConnections    int    `toml:"MaxOpenConnections"`
	MaxIdleConnections    int    `toml:"MaxIdleConnections"`
	ConnectionMaxLifetime time.Duration
	ConnectionMaxIdleTime time.Duration
	DebugMode             bool `json:"DebugMode"`
}
