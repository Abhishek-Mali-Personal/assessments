package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"log"
	"os"
)

var (
	DriverName = "postgres"
	DB         *sql.DB
)

func InitDatabase() {
	var (
		dsn string
		err error
	)
	DriverName = os.Getenv("DRIVER")
	if DriverName == "mysql" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			os.Getenv("USER"), os.Getenv("PASSWORD"),
			os.Getenv("HOST"), os.Getenv("PORT"),
			os.Getenv("DBNAME"))
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", os.Getenv("HOST"), os.Getenv("USER"),
			os.Getenv("PASSWORD"), os.Getenv("DBNAME"), os.Getenv("PORT"), "disable")
	}
	DB, err = sql.Open(DriverName, dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
}
