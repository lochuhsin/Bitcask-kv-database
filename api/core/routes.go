package core

import "github.com/gin-gonic/gin"

func Routes(route *gin.Engine) {
	coreRoute := route.Group("/core")
	coreRoute.GET("/", getHandler)
	coreRoute.POST("/", postHandler)
	coreRoute.PATCH("/", patchHandler)
	coreRoute.DELETE("/", deleteHandler)
	coreRoute.POST("/", watchHandler)
	coreRoute.POST("/", postSyncHandler)
	coreRoute.PATCH("/", patchSyncHandler)
	coreRoute.DELETE("/", deleteSyncHandler)
}
