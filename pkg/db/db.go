// Package db contains specific primitives for database configuration and connections.
package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gaganchawara/loans/internal/errorcode"
	"github.com/gaganchawara/loans/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewDB initializes a new database connection based on the provided configuration.
func NewDB(ctx context.Context, c Config) (*gorm.DB, errors.Error) {
	dialect := mysql.Open(fmt.Sprintf(MysqlConnectionDSNFormat, c.Username, c.Password, c.Protocol, c.URL, c.Name))

	db, err := gorm.Open(dialect, &gorm.Config{
		AllowGlobalUpdate:      false,
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		// Set log level based on debug mode
		Logger: logger.Default.LogMode(getLogLevelByDebugMode(c.DebugMode)),
	})
	if err != nil {
		return nil, errors.New(ctx, errorcode.InternalServerError, err).Report()
	}

	// Configure database connection parameters.
	var dbConn *sql.DB
	if dbConn, err = db.DB(); err != nil {
		return nil, errors.New(ctx, errorcode.InternalServerError, err).Report()
	}

	dbConn.SetMaxIdleConns(c.MaxIdleConnections)
	dbConn.SetMaxOpenConns(c.MaxOpenConnections)
	dbConn.SetConnMaxLifetime(c.ConnectionMaxLifetime * time.Second)
	dbConn.SetConnMaxIdleTime(c.ConnectionMaxIdleTime * time.Second)

	return db, nil
}

// getLogLevelByDebugMode returns the logger log level based on the debug mode.
// If the application's database is in debug mode, it sets the log level to info.
// The default log level for GORM DB is warning, which can be overridden by this method.
func getLogLevelByDebugMode(debug bool) logger.LogLevel {
	if debug == false {
		return logger.Silent
	} else {
		return logger.Info
	}
}
