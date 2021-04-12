package routes

import (
	"gin_psql_microservice_template/configs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// GetRoutes is a function which returns the information of registered routes
func GetRoutes() []RouteInfo {
	// controller := controllers.NewController()
	routeInfo := []RouteInfo{}
	return routeInfo
}

// SetupRoutes sets the routes
func SetupRoutes(engine *gin.Engine, models []RouteInfo) {
	ENABLE_SWAGGER := configs.Envs["ENABLE_SWAGGER"]
	if ENABLE_SWAGGER == "true" {
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	for _, controller := range models {
		engine.Handle(controller.Method, controller.Path, controller.Handler)
	} // setup routes
}
