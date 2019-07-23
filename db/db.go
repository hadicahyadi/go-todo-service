package db

import (
  "database/sql"
  "fmt"
  "log"
  "os"

  _ "github.com/go-sql-driver/mysql"
  "github.com/joho/godotenv"
)

var db *sql.DB
var err error

func Connect() *sql.DB {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
  dbHost := os.Getenv("DB_HOST")
  dbUser := os.Getenv("DB_USER")
  dbPass := os.Getenv("DB_PASSWORD")
  dbName := os.Getenv("DB_NAME")
  dbPort := os.Getenv("DB_PORT")

  dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName)
  
  db, err = sql.Open("mysql", dbUrl)
  if err != nil {
		log.Fatal(err)
	} else {
    fmt.Println("Database connected.")
  }
  return db
}
