package cluster

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary cluster status
// @Schemes http
// @Description cluster status
// @Success 200 {object} ClusterStatusSchema
// @Router /cluster/status [get]
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

// @Summary cluster configuration
// @Schemes http
// @Description cluster configuration
// @Success 200 {object} ClusterConfigurationSchema
// @Router /cluster/configuration [get]
func getConfigHandler(c *gin.Context) {
	/**
	 * Retrieves the cluster configuration
	 */
	c.JSON(http.StatusOK, &ClusterConfigurationSchema{10})
}
