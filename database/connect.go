package database

import (
	"api-fiber-gorm/config"
	"api-fiber-gorm/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strconv"
)

var BookDB *gorm.DB
var PostDB *gorm.DB

func ConnectBookDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic(err)
	}

	BookDB, err = gorm.Open(postgres.Open(fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"),
		port,
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
	)))
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	err = BookDB.AutoMigrate(&models.Product{})
	if err != nil {
		return
	}

	fmt.Println("Database Migrated")
}

func ConnectPostDB() {
	var err error
	dsn := "host=localhost port=5432 user=postgres password=Root dbname=posts_db sslmode=disable"
	PostDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}
