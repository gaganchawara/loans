package boot

import (
	"context"
	"github.com/gaganchawara/loans/pkg/db"
	"gorm.io/gorm"
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
	DB *gorm.DB
)

func Initialize(ctx context.Context) {
	var err error
	var ierr errors.Error

	logger.Get(ctx).Info("booting the application")

	// initializes error packages with Hooks
	errors.Initialize(logger.ErrorLogger())
	err = pkgconfig.LoadConfig(getConfigPath(), GetEnv(), &cfg)
	if err != nil {
		panic(err)
	}

	DB, ierr = db.NewDB(ctx, cfg.DB)
	if ierr != nil {
		panic(ierr)
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
