package bootstrap

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerHandler(c *gin.Context) {
	c.Bind(&registerSchema{})
	/**
	 * Register the rebitcask components to the cluster
	 */
}

func retrievePeersHandler(c *gin.Context) {
	/**
	 * Retrieving the list of all existing members
	 * in the cluster
	 */
	c.JSON(http.StatusOK, peerListSchema{})
}
