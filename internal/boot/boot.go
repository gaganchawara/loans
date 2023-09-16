package boot

import (
	"context"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gaganchawara/loans/internal/config"
	pkgconfig "github.com/gaganchawara/loans/pkg/config"
	"github.com/gaganchawara/loans/pkg/errors"
	"github.com/gaganchawara/loans/pkg/logger"
)

var (
	cfg config.Config
)

func Initialize(ctx context.Context) {
	logger.Get(ctx).Info("booting the application")

	// initializes error packages with Hooks
	errors.Initialize(logger.ErrorLogger())
	err := pkgconfig.LoadConfig(getConfigPath(), GetEnv(), &cfg)
	if err != nil {
		panic(err)
	}
}

func getConfigPath() string {
	return filepath.Join(getRootPath(), "config")
}

func getRootPath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../..")
}

func GetEnv() string {
	environment := os.Getenv("APP_ENV")
	if environment == "" {
		environment = "default"
	}

	return environment
}
