package main

import (
	"rebitcask/discovery/api/bootstrap"
	"rebitcask/discovery/api/cluster"
	_ "rebitcask/discovery/docs"
	"rebitcask/discovery/settings"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	/**
	 * Setting up server as
	 * http://localhost:8765/_rebitcask/........
	 */
	Init()
	r := gin.Default()
	r.Group("/_rebitcask") // main prefix
	bootstrap.Routes(r)
	cluster.Routes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(settings.Config.HTTP_PORT)
}
