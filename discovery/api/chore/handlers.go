package chore

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthzHandler(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
