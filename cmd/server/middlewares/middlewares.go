package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

// Setup is a function for setting up middlewares
func Setup(engine *gin.Engine) {
	logger := logrus.New()
	engine.Use(gin.Recovery())
	engine.Use(ginlogrus.Logger(logger))
	engine.Use(cors.Default())
}
