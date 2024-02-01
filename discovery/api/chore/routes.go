package chore

import "github.com/gin-gonic/gin"

func Routes(route *gin.Engine) {
	route.GET("/healthz", healthzHandler)
}
