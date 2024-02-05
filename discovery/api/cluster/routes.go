package cluster

import "github.com/gin-gonic/gin"

func Routes(route *gin.Engine) {
	r := route.Group("/cluster")
	r.GET("/status", getStatusHandler)
	r.GET("/configuration", getConfigHandler)
	r.POST("/register/", registerHandler)
	r.GET("/peers", retrievePeersHandler)
	r.POST("/finished-peer/", finishedPeerHandler)
}
