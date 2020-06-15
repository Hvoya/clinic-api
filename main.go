package main

import (
	"clinic/controller"
	"clinic/database"
	"clinic/model"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	db := database.Init()
	model.Migrate(db)

	controller.InitControllers(db)
}
