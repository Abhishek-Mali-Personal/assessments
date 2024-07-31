package main

import (
	"github.com/Abhishek-Mali-Simform/assessments/database"
	"github.com/Abhishek-Mali-Simform/assessments/routers"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env file not found")
	}
	routers.InitRoute()
	database.InitDatabase()
}

func main() {
	err := routers.Route.Run(":5000")
	if err != nil {
		log.Fatal("unable to start server")
	}
}
