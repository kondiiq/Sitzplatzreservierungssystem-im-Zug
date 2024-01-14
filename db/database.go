package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host = ""
	port = 2137
	user = ""
	password
	dbName = ""
)

func DbConnection() *gorm.DB {

	sqlInfo := fmt.Sprintf("host=%s \n port=%s \n user=%s \n password=%s \n dbname=%s sslmode=disabled",
		host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	panic(err)

	return db
}
