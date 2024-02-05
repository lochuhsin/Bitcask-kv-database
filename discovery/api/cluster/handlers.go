package cluster

import (
	"net/http"
	"rebitcask/discovery/cache"
	"rebitcask/discovery/settings"

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
	status := ClusterStatus(cache.ClusterCache.Get(c.Request.Context(), cache.Status))
	c.JSON(http.StatusOK, &ClusterStatusSchema{Status: status})
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

// @BasePath /api/v1
// @Summary register cluster members
// @Schemes http
// @Description register cluster members
// @Param RequestBody body registerRequestSchema true "register cluster members"
// @Success 200 {object} registerResponseSchema
// @Router /cluster/register [post]
func registerHandler(c *gin.Context) {
	obj := registerRequestSchema{}
	c.Bind(&obj)
	/**
	 * Register the rebitcask components to the cluster
	 */
	status := cache.PeerCache.Add(c.Request.Context(), cache.PeerCacheSchema(obj))
	if !status {
		c.JSON(http.StatusBadRequest, registerResponseSchema{
			Message: "Invalid operation, the seats were full",
		})
		return
	}

	// TODO: definitely a bug ...
	// since we are doing two operations in concurrency programming lol
	if cache.PeerCache.Count(c.Request.Context()) == settings.Config.CLUSTER_MEMBER_COUNT {
		cache.ClusterCache.Set(c.Request.Context(), cache.Status, string(YELLO))
	}

	c.JSON(http.StatusAccepted, registerResponseSchema{
		Message: "ok",
	})
}

// @Summary get all registered cluster members
// @Schemes http
// @Description get all cluster members
// @Success 200 {object} peerListResponseSchema
// @Router /cluster/peers [get]
func retrievePeersHandler(c *gin.Context) {
	/**
	 * Retrieving the list of all existing members
	 * in the cluster
	 */
	peers := cache.PeerCache.GetAll(c.Request.Context())
	peerResponses := make([]peerSchema, len(peers))
	for i, p := range peers {
		peerResponses[i] = peerSchema(p)
	}

	c.JSON(http.StatusOK, peerListResponseSchema{
		Peers: peerResponses,
	})
}

// @Summary get all registered cluster members
// @Schemes http
// @Description get all cluster members
// @Success 200 {object} peerSchema
// @Param RequestBody body peerSchema true "when the peer finished everything, waiting cluster to startup call this api"
// @Router /cluster/finished-peer/ [post]
func finishedPeerHandler(c *gin.Context) {
	context := c.Request.Context()
	counter := cache.CounterCache

	// TODO: definitely a bug ...
	// since we are doing two operations in concurrency programming lol
	counter.Add(context)
	if counter.Count(context) >= settings.Config.CLUSTER_MEMBER_COUNT {
		cache.ClusterCache.Set(context, cache.Status, string(GREEN))
	}
}
