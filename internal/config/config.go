package config

import "github.com/gaganchawara/loans/pkg/db"

type Config struct {
	App App
	DB db.Config
}

type App struct {
	Env             string `toml:"Env"`
	ServiceName     string `toml:"ServiceName"`
	Hostname        string `toml:"Hostname"`
	Port            string `toml:"Port"`
	ShutdownTimeout int    `toml:"ShutdownTimeout"`
	ShutdownDelay   int    `toml:"ShutdownDelay"`
	GitCommitHash   string `toml:"GitCommitHash"`
}