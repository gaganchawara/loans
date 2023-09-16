// Package db has specific primitives for database config & connections.
//
// Usage:
// - 	E.g. dbpkg.NewDb(&c), where c must implement ConfigReader and default use case is to just use Config struct.
package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gaganchawara/loans/internal/errorcode"
	"github.com/gaganchawara/loans/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

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

// getLogLevelByDebugMode return logger log level based on debug mode.
// If app db is in debug mode, make log level as info
// Default log level for gorm db is warning, overriding that by this method.
func getLogLevelByDebugMode(debug bool) logger.LogLevel {
	if debug == false {
		return logger.Silent
	} else {
		return logger.Info
	}
}
