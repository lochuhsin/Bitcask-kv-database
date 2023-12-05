package core

import (
	"log"
	"net/http"

	"rebitcask/internal"

	"github.com/gin-gonic/gin"
)

func getHandler(c *gin.Context) {
	key := c.Query("key")
	val, status := internal.Get(key)
	if !status {
		c.String(http.StatusBadRequest, "")
	}
	c.JSON(http.StatusOK, val)
}

func postHandler(c *gin.Context) {
	obj := dataRequestSerializer{}
	c.Bind(&obj)
	internal.Set(obj.Key, obj.Value)
	log.Println(obj.Key, obj.Value)
	c.JSON(http.StatusCreated, obj)
}

func putHandler(c *gin.Context) {
	obj := dataRequestSerializer{}
	c.Bind(&obj)
	err := internal.Set(obj.Key, obj.Value)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid operation")
	}
	c.JSON(http.StatusNoContent, "")
}

func patchHandler(c *gin.Context) {
	obj := dataRequestSerializer{}
	c.Bind(&obj)

	err := internal.Set(obj.Key, obj.Value)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid operation")
	}
	c.JSON(http.StatusNoContent, "")
}

func deleteHandler(c *gin.Context) {
	obj := dataRequestSerializer{}
	c.Bind(&obj)
	err := internal.Delete(obj.Key)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid operation")
	}
	c.JSON(http.StatusAccepted, "")
}
