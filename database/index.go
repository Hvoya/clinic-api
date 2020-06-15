package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

func Init() *gorm.DB {
	db_name, _ := os.LookupEnv("DB_NAME")
	db_user, _ := os.LookupEnv("DB_USER")
	db_password, _ := os.LookupEnv("DB_PASSWORD")

	dbString := fmt.Sprintf("host=localhost port=5432 dbname=%v user=%v password=%v sslmode=disable", db_name, db_user, db_password)

	db, _ := gorm.Open("postgres", dbString)
	return db
}
