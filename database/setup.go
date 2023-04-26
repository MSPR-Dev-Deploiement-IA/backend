package database

import (
	"backend/models"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	var db *gorm.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}

		fmt.Printf("Failed to connect to database. Retrying in %d seconds...\n", i*5)
		time.Sleep(time.Duration(i*5) * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after 10 attempts: %w", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Printf("Failed to migrate database schema: %v\n", err)
		return nil, err
	}

	fmt.Println("Successfully migrated database schema")
	return db, nil
}
