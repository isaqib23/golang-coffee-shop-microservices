package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type DBClient struct {
	DB *sql.DB
}

func NewDatabaseConnection() *DBClient {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	dbHostName := os.Getenv("DB_HOSTNAME")
	dbPort := os.Getenv("DB_PORT")

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@("+dbHostName+":"+dbPort+")/"+dbName)
	if err != nil {
		fmt.Println("DB Error")
		panic(err.Error())
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &DBClient{db}
}
