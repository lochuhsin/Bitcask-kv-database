package bootstrap

import "github.com/gin-gonic/gin"

func Routes(route *gin.Engine) {
	r := route.Group("/bootstrap")
	r.POST("/register", registerHandler)
	r.GET("/peers", retrievePeersHandler)
}
