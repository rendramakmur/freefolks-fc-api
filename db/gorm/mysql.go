package gorm

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMySQL() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbCharset := os.Getenv("DB_CHARSET")
	dbLocal := os.Getenv("DB_LOCAL")

	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=%s&loc=%s&parseTime=True", dbUser, dbPassword, dbHost, dbName, dbCharset, dbLocal)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	return db
}
