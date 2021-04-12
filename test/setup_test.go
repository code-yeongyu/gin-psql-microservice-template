package test

import (
	"gin_psql_microservice_template/cmd/server/middlewares"
	"gin_psql_microservice_template/cmd/server/models"
	"gin_psql_microservice_template/cmd/server/routes"
	"gin_psql_microservice_template/cmd/server/utils"
	"testing"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func TestEngineSetup(t *testing.T) {
	engine = gin.New()
	utils.InitDB(models.GetModels())
	middlewares.Setup(engine)
	routes.SetupRoutes(engine, routes.GetRoutes())
}
