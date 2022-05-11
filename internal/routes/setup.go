package routes

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// SetupRoutes connects the HTTP API endpoints to the handlers
func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./website", true)))

	return r
}
