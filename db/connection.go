package db

import (
	models "Plane-Booking/db/Models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	var err error

	host := "localhost"
	databaseName := "plane"
	port := "5432"

	dsn := fmt.Sprintf("host=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		host, databaseName, port)

	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
}

func Migrate() {
	Database.AutoMigrate(&models.User{})
	Database.AutoMigrate(&models.Plane{})
	Database.AutoMigrate(&models.Seat{})
	Database.AutoMigrate(&models.Booking{})
}
