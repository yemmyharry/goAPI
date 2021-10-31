package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func Dbsetup() *gorm.DB {
	errN := godotenv.Load()
	if errN != nil {
		log.Fatal("failed to load env file")
	}

	dbuser := os.Getenv("DBUSER")
	dbpass := os.Getenv("DBPASS")
	dbhost := os.Getenv("DBHOST")
	dbname := os.Getenv("DBNAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpass, dbhost, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to create db connection")
	}

	return db

}

func CloseDBConnection(db *gorm.DB) {
	dbsql, errN := db.DB()
	if errN != nil {
		log.Fatal("failed to close connection from database")
	}
	err := dbsql.Close()
	if err != nil {
		return
	}
}
