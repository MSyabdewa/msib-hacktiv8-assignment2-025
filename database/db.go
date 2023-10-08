package database

import (
	"fmt"

	"github.com/MSyabdewa/msib-hacktiv8-assignment2-025/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DBHost     = "localhost"
	DBPort     = 5432
	DBUser     = "dew"
	DBPassword = "root"
	DBName     = "assignment2"
)

func InitializeDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DBHost, DBPort, DBUser, DBPassword, DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.Debug().AutoMigrate(&models.Order{}, &models.Item{})

	return db, nil
}
