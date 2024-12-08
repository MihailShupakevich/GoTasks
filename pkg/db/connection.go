package db

import (
	"awesomeProject/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connection() *gorm.DB {
	dsn := "host=localhost user=admin dbname=GoDB password=admin sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate([]domain.Task{})
	return db
}
