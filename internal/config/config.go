package config

import (
	"github.com/gaganchawara/loans/pkg/db"
	"github.com/gaganchawara/loans/pkg/grpcserver"
)

type Config struct {
	App App
	DB  db.Config
}

type App struct {
	Env             string                     `toml:"Env"`
	ServiceName     string                     `toml:"ServiceName"`
	ServerAddresses grpcserver.ServerAddresses `json:"ServerAddresses"`
	ShutdownTimeout int                        `toml:"ShutdownTimeout"`
	ShutdownDelay   int                        `toml:"ShutdownDelay"`
	GitCommitHash   string                     `toml:"GitCommitHash"`
}
