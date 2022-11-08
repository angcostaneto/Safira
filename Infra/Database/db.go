package Database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(
		postgres.Open("host=localhost user=test password=test dbname=test port=5432 sslmode=disable"),
		&gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}
