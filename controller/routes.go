package controller

import (
	"github.com/gin-gonic/gin"
)

func addRoutes(r *gin.Engine) {
	// add routes here
	r.GET("hello", func(ctx *gin.Context) {
		ctx.String(200, "hello")
	})
}
