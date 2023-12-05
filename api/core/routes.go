package core

import "github.com/gin-gonic/gin"

func Routes(route *gin.Engine) {
	coreRoute := route.Group("/core")
	coreRoute.GET("/", getHandler)
	coreRoute.POST("/", postHandler)
	coreRoute.PATCH("/", patchHandler)
	coreRoute.PUT("/", putHandler)
	coreRoute.DELETE("/", deleteHandler)
}
