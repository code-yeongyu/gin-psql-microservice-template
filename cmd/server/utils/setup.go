package utils

import (
	"gin_psql_microservice_template/cmd/server/docs"
	"gin_psql_microservice_template/cmd/server/routes"
	"gin_psql_microservice_template/configs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRoutes sets the routes as setted in routes/routes.go
func SetupRoutes(engine *gin.Engine) {
	ENABLE_SWAGGER := configs.Envs["ENABLE_SWAGGER"]
	if ENABLE_SWAGGER == "true" {
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	for _, controller := range routes.GetRoutes() {
		engine.Handle(controller.Method, controller.Path, controller.Handler)
	} // setup routes
}

// SetupDocuments
func SetupDocuments() {
	docs.SwaggerInfo.Title = "API Document"
	docs.SwaggerInfo.Description = "The REST API Document"
	docs.SwaggerInfo.Host = "localhost"
	docs.SwaggerInfo.Version = "0.1"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}
}
