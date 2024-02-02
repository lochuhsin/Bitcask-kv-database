package chore

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary root path handler
// @Schemes http
// @Description root path handler
// @Accept json
// @Produce json
// @Success 200 {object} rootResponseSchema
// @Router / [get]
func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, rootResponseSchema{
		Message: "Welcome to rebitcask",
	})
}

// @Summary healthz check endpoint
// @Schemes http
// @Description healthz check endpoint
// @Accept json
// @Produce json
// @Success 200 {object} healthzResponseSchema
// @Router /healthz [get]
func healthzHandler(c *gin.Context) {
	c.JSON(http.StatusOK, healthzResponseSchema{
		Message: "ok",
	})
}
