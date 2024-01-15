package core

import (
	"log"
	"net/http"
	"rebitcask"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /core [get]
func getHandler(c *gin.Context) {
	key := c.Query("key")
	val, status := rebitcask.Get(key)
	if !status {
		c.String(http.StatusBadRequest, "")
	}
	c.JSON(http.StatusOK, val)
}

func postHandler(c *gin.Context) {
	obj := dataRequestSerializer{}
	c.Bind(&obj)
	rebitcask.Set(obj.Key, obj.Value)
	log.Println(obj.Key, obj.Value)
	c.JSON(http.StatusCreated, obj)
}

func putHandler(c *gin.Context) {
	obj := dataRequestSerializer{}
	c.Bind(&obj)
	err := rebitcask.Set(obj.Key, obj.Value)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid operation")
	}
	c.JSON(http.StatusNoContent, "")
}

func patchHandler(c *gin.Context) {
	obj := dataRequestSerializer{}
	c.Bind(&obj)

	err := rebitcask.Set(obj.Key, obj.Value)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid operation")
	}
	c.JSON(http.StatusNoContent, "")
}

func deleteHandler(c *gin.Context) {
	obj := dataRequestSerializer{}
	c.Bind(&obj)
	err := rebitcask.Delete(obj.Key)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid operation")
	}
	c.JSON(http.StatusAccepted, "")
}
