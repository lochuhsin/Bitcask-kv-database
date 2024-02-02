package chore

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary healthz endpoint
// @Schemes http
// @Description healthz endpoint
// @Success 200 {object} healthzResponseSchema
// @Router /healthz [get]
func healthzHandler(c *gin.Context) {
	c.JSON(http.StatusOK, healthzResponseSchema{
		Message: "ok",
	})
}
