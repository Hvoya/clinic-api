package controller

import (
	"clinic/model"
	"github.com/jinzhu/gorm"
	"net/http"
)

func RegisterUser(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := model.Direction{Name: "134"}
		db.Create(&user)
	}
}
