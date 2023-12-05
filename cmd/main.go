package main

import (
	"rebitcask/api/chore"
	"rebitcask/api/core"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	core.Routes(r)
	chore.Route(r)
	port := ":8000"
	r.Run(port) // listen and serve on 0.0.0.0:8080

	// TODO: Add graceful shutdown
}
