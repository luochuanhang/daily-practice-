package storage

import (
	"errors"

	"gorm.io/gorm"
)

var (
	// ErrNil stands for connection nil error.
	ErrNil = errors.New("storage connection is nil")
)

// Storage Storage
type Storage interface {
	Name() string
	Init(url string) error
	Get() *gorm.DB
	New(url string) (*gorm.DB, error)
}

// NewStorage creates a new storage.
func NewStorage(typ string) Storage {
	switch typ {
	case "mysql":
		return &mysql{}
	default:
		return defaultmysql
	}
}
