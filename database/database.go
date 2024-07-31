package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DB *sql.DB

func InitDatabase() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", os.Getenv("HOST"), os.Getenv("USER"),
		os.Getenv("PASSWORD"), os.Getenv("DBNAME"), os.Getenv("PORT"), "disable")
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
}
