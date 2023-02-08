package router

import "github.com/gin-gonic/gin"

func Start() {
	// creates an engine
	// Creates a gin router with default middleware - logger and recovery
	routerEngine := gin.Default()
	routerEngine.BasePath("/api")
}
