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
		return nil
	}

	err = db.AutoMigrate(
		&entity.Project{},
		&entity.Team{},
		&entity.User{},
	)

	if err != nil {
		return nil
	}

	fmt.Println("Connected to database")

	return db
}
