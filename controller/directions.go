package controller

import (
	"clinic/model"
	"clinic/utils"
	"github.com/jinzhu/gorm"
	"net/http"
)

// GetDirections godoc
// @Summary Get all directions
// @Tags directions
// @Router /directions [get]
// @Success 200 {array} model.Direction
func GetDirections(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var directions []model.Direction
		db.Find(&directions)

		utils.SendJSON(w, directions)
	}
}
