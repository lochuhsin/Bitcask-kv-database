package bootstrap

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary register cluster members
// @Schemes http
// @Description register cluster members
// @Param RequestBody body registerRequestSchema true "register cluster members"
// @Success 200 {object} registerResponseSchema
// @Router /bootstrap/register [post]
func registerHandler(c *gin.Context) {
	c.Bind(&registerRequestSchema{})
	/**
	 * Register the rebitcask components to the cluster
	 */
	c.JSON(http.StatusAccepted, registerResponseSchema{
		Message: "ok",
	})
}

// @Summary get all registered cluster members
// @Schemes http
// @Description get all cluster members
// @Success 200 {object} peerListResponseSchema
// @Router /bootstrap/peers [get]
func retrievePeersHandler(c *gin.Context) {
	/**
	 * Retrieving the list of all existing members
	 * in the cluster
	 */
	c.JSON(http.StatusOK, peerListResponseSchema{})
}
