package test

import (
	"gin_psql_microservice_template/cmd/server/middlewares"
	"gin_psql_microservice_template/cmd/server/utils"
	"testing"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func TestEngineSetup(t *testing.T) {
	engine = gin.New()
	utils.InitDB()
	middlewares.Setup(engine)
	utils.SetupRoutes(engine)
}
