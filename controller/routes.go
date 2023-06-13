package controller

import (
	"common-server/controller/handlers"
	"github.com/gin-gonic/gin"
)

func addRoutes(r *gin.Engine) {
	// add routes here
	r.GET("hello", handlers.Hello)
}
