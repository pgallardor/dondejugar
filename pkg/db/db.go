package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDbStorage(cfg postgres.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	return db, err
}
