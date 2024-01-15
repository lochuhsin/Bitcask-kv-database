package chore

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	route.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})
	route.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, world")
	})
}
