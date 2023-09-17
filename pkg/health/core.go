package health

import (
	"context"
	"fmt"
	"sync"

	"gorm.io/gorm"
)

// Core holds business logic and/or orchestrator of other things in the package.
type Core struct {
	isHealthy bool
	mutex     sync.RWMutex
	db        *gorm.DB
}

// NewCore creates Core.
func NewCore(db *gorm.DB) *Core {
	return &Core{
		isHealthy: true,
		db:        db,
	}
}

// RunHealthCheck runs various server checks and returns true if all individual components are working fine.
func (c *Core) RunHealthCheck(ctx context.Context) error {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if !c.isHealthy {
		return fmt.Errorf("server marked unhealthy")
	}

	sqlDB, err := c.db.DB()
	if err != nil {
		return err
	}

	if err = sqlDB.PingContext(ctx); err != nil {
		return err
	}

	return nil
}

// Ping checks whether the app is able to receive requests
func (c *Core) Ping(ctx context.Context) error {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if !c.isHealthy {
		return fmt.Errorf("server marked unhealthy")
	}

	return nil
}
