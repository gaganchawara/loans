package config

import (
	"github.com/gaganchawara/loans/pkg/db"
	"github.com/gaganchawara/loans/pkg/grpcserver"
	interceptors "github.com/gaganchawara/loans/pkg/grpcserver/interceptor"
)

type Config struct {
	App  App
	Auth interceptors.BasicAuthCreds
	DB   db.Config
}

type App struct {
	Env             string                     `toml:"Env"`
	ServiceName     string                     `toml:"ServiceName"`
	ServerAddresses grpcserver.ServerAddresses `json:"ServerAddresses"`
	ShutdownTimeout int                        `toml:"ShutdownTimeout"`
	ShutdownDelay   int                        `toml:"ShutdownDelay"`
	GitCommitHash   string                     `toml:"GitCommitHash"`
}
