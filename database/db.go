package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

const ()

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot load env file :(")
	}
}

func databaseConnection() (*gorm.DB, error) {
	loadEnv()
	host := os.Getenv("HOST_NAME")
	user := os.Getenv("USER")
	passwd := os.Getenv("DB_PASSWD")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DBNAME")
	dsn := fmt.Sprintf("host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable", host, user, passwd, port, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	return db, nil
}
