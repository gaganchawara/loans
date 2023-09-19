package entity

import (
	"database/sql"
	"time"
)

type Entity struct {
	CreatedAt sql.NullTime `gorm:"column:created_at" json:"created_at"`
	UpdatedAt sql.NullTime `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
}

func (e *Entity) RefreshTimestamps() {
	now := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	if !e.CreatedAt.Valid {
		e.CreatedAt = now
	}
	e.UpdatedAt = now
}
