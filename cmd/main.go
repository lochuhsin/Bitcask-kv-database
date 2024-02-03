package main

import (
	"rebitcask"
	"rebitcask/api/chore"
	"rebitcask/api/core"

	_ "rebitcask/docs"

	"rebitcask/internal/settings"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func main() {
	flags := ParseFlags()
	rebitcask.Init(flags.envPaths...)

	if settings.Config.MODE == settings.CLUSTER {
		func() {
			/**
			 * 1. register to discvery server
			 * 2. wait the server status to become yellow
			 * 3. retrive all existing peerlist from server
			 * 4. request back to server that everything is ok
			 * 5. wait the server to become green
			 * 6. start running raft ...
			 */

		}()
	}

	env := settings.Config
	r := gin.Default()
	core.Routes(r)
	chore.Routes(r)

	// starts swagger at localhost:port/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	go runGRPC(env.GRPC_PORT)
	r.Run(env.HTTP_PORT) // listen and serve on 0.0.0.0:8080

}
