package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := "postgres://demystif:gppw2023@localhost:5432/gin-postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
