package main

import (
	"flag"
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
	flag.Var(&envPaths, "envfiles", "Specifies the env files")
	flag.Parse()
	rebitcask.Init([]string(envPaths)...)

	env := settings.Config
	r := gin.Default()
	core.Routes(r)
	chore.Routes(r)

	// starts swagger at localhost:port/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	go runGRPC(env.GRPC_PORT)
	r.Run(env.HTTP_PORT) // listen and serve on 0.0.0.0:8080

}
