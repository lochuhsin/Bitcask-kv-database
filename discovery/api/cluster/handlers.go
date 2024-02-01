package cluster

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getStatusHandler(c *gin.Context) {
	/**
	 * Responses the cluster status
	 * borrowing from elasticsearch
	 * 1. Green
	 * 2. Yellow
	 * 3. Red
	 */

	c.JSON(http.StatusOK, &ClusterStatusSchema{Status: GREEN})
}

func getConfigHandler(c *gin.Context) {
	/**
	 * Retrieves the cluster configuration
	 */
	c.JSON(http.StatusOK, &ClusterConfigrationSchema{10})
}
