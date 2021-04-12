package main

import (
	"gin_psql_microservice_template/cmd/server/middlewares"
	"gin_psql_microservice_template/cmd/server/models"
	"gin_psql_microservice_template/cmd/server/routes"
	"gin_psql_microservice_template/cmd/server/utils"
	"gin_psql_microservice_template/configs"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func runServer(engine *gin.Engine) {
	PORT := ":" + configs.Envs["SERVER_PORT"]

	engine.Run(PORT)
}

func main() {
	govalidator.SetFieldsRequiredByDefault(false)
	utils.InitDB(models.GetModels())
	engine := gin.New()
	middlewares.Setup(engine)
	utils.SetupDocuments()
	routes.SetupRoutes(engine, routes.GetRoutes())
	runServer(engine)
}
