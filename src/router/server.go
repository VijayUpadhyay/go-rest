package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
	"learn/go-rest/src/config"
)

func Start() {
	// creates an engine
	// Creates a gin router with default middleware - logger and recovery
	routerEngine := gin.Default()
	if logger, ok := config.InitLogger(); ok {
		routerEngine.Use(ginlogrus.Logger(logger))
	}

	//routerEngine.Handlers
	routerEngine.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "Application is up!!"})
	})
	err := routerEngine.Run(":9090")
	if err != nil {
		logrus.Errorf("Applicati start up failed: %+v", err)
		return
	}
}
