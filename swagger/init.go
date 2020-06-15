package swagger

import (
	"clinic/docs"
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

func InitSwagger(router *chi.Mux) {
	docs.SwaggerInfo.Title = "Clinic Api"
	docs.SwaggerInfo.Description = "Simple api for clinic app"
	docs.SwaggerInfo.Version = "1.0"
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:1323/swagger/doc.json"),
	))
}
