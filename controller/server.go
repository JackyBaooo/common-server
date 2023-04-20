package controller

import (
	"common-server/config"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartServer() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	addRoutes(r)
	host := config.GlobalViper.GetString("server.host")
	port := config.GlobalViper.GetString("server.port")
	addr := fmt.Sprintf("%s:%s", host, port)
	log.Infof("server will run on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
