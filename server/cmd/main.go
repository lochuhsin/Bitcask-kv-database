package main

import (
	"rebitcask"
	"rebitcask/server/api/chore"
	"rebitcask/server/api/core"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func main() {
	rebitcask.Init()

	r := gin.Default()
	core.Routes(r)
	chore.Route(r)
	port := ":8000"
	ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.URL("http://localhost:8000/swagger/doc.json"))
	r.Run(port) // listen and serve on 0.0.0.0:8080

	// TODO: Add graceful shutdown
}
