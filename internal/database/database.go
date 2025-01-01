package database

import (
	"fmt"
	"os"

	"github.com/IlhamRanggaKurniawan/Teamers.git/internal/database/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Something went wrong inside the database")
	}

	err = db.AutoMigrate(
		entity.Class{},
		entity.User{},
		entity.Task{},
		entity.Subject{},
		entity.Schedule{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Failed to migrate the database")
	}

	fmt.Println("Connected to database")

	return db
}
