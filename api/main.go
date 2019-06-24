package main

import (
	"app/cmd"
	"app/database"
	"app/router"
	"log"

	"github.com/joho/godotenv"
)

func init() {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	database.Load()
	router.Load()
}

func main() {
	cmd.Execute()
}
