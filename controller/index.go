package controller

import (
	"clinic/swagger"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"net/http"
)

func InitControllers(db *gorm.DB) {
	router := chi.NewRouter()
	swagger.InitSwagger(router)

	router.Get("/", RegisterUser(db))
	router.Get("/directions", GetDirections(db))
	err := http.ListenAndServe(":1323", router)

	if err != nil {
		panic(err)
	}
}
