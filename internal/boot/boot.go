package boot

import (
	"context"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gaganchawara/loans/pkg/tracing"

	"github.com/gaganchawara/loans/internal/errorcode"

	"github.com/dlmiddlecote/sqlstats"
	"github.com/gaganchawara/loans/pkg/db"
	"github.com/prometheus/client_golang/prometheus"
	"gorm.io/gorm"

	"github.com/gaganchawara/loans/internal/config"
	pkgconfig "github.com/gaganchawara/loans/pkg/config"
	"github.com/gaganchawara/loans/pkg/errors"
	"github.com/gaganchawara/loans/pkg/logger"
)

var (
	// Config represents the application configuration.
	Config config.Config
	// DB is the database connection.
	DB *gorm.DB
)

// Initialize serves as the universal bootstrapping function, responsible for common
// initialization routines across various parts of the application.
func Initialize(ctx context.Context) errors.Error {
	var err error
	var ierr errors.Error

	logger.Get(ctx).Info("booting the application")

	// initializes error packages with Hooks
	errors.Initialize(logger.ErrorLogger())

	err = pkgconfig.LoadConfig(getConfigPath(), GetEnv(), &Config)
	if err != nil {
		return errors.New(ctx, errorcode.InternalServerError, err).Report()
	}

	Config.Tracing.ServiceName = Config.App.ServiceName
	Config.Tracing.Env = Config.App.Env

	// Initialize a new database connection.
	DB, ierr = db.NewDB(ctx, Config.DB)
	if ierr != nil {
		return errors.New(ctx, errorcode.InternalServerError, err).Report()
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return errors.New(ctx, errorcode.InternalServerError, err).Report()
	}

	// Register a SQL database statistics collector for Prometheus.
	collector := sqlstats.NewStatsCollector(Config.DB.URL+"-"+Config.DB.Name, sqlDB)
	prometheus.MustRegister(collector)

	if ierr = tracing.InitTracer(ctx, Config.Tracing, logger.Get(ctx)); ierr != nil {
		return ierr
	}

	return nil
}

// getConfigPath returns the path to the configuration directory
func getConfigPath() string {
	return filepath.Join(getRootPath(), "config")
}

// getRootPath returns the root directory of the application
func getRootPath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../..")
}

// GetEnv retrieves the application's runtime environment.
func GetEnv() string {
	environment := os.Getenv("APP_ENV")
	if environment == "" {
		environment = "default"
	}

	return environment
}
