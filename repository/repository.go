// This file contains the repository implementation layer.
package repository

import (
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}
