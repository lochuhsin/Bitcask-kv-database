package core

import (
	"log"
	"net/http"
	"rebitcask"

	"github.com/gin-gonic/gin"
)

/**
 * The table of swagger apis
 * https://github.com/swaggo/swag/blob/master/README_zh-CN.md#api%E6%93%8D%E4%BD%9C
 */

// @BasePath /api/v1

// @Summary get value by key
// @Schemes http
// @Description get value by key
// @Param key query string true "query database with key"
// @Success 200 {string} string
// @Router /core [get]
func getHandler(c *gin.Context) {
	key := c.Query("key")
	val, status := rebitcask.Get(key)
	if !status {
		c.String(http.StatusBadRequest, "")
	}
	c.JSON(http.StatusOK, val)
}

// @Summary insert key / value
// @Schemes http
// @Description insert key / value
// @Accept json
// @Produce json
// @Param RequestBody body dataRequestSchema true "request body for create an entry"
// @Success 200 {object} dataRequestSchema
// @Router /core [post]
func postHandler(c *gin.Context) {
	obj := dataRequestSchema{}
	c.Bind(&obj)
	rebitcask.Set(obj.Key, obj.Value)
	log.Println(obj.Key, obj.Value)
	c.JSON(http.StatusCreated, obj)
}

// @Summary update key / value
// @Schemes http
// @Description update key / value
// @Accept json
// @Produce json
// @Param RequestBody body dataRequestSchema true "request body for update an entry"
// @Success 200 {object} dataRequestSchema
// @Router /core [patch]
func patchHandler(c *gin.Context) {
	obj := dataRequestSchema{}
	c.Bind(&obj)

	err := rebitcask.Set(obj.Key, obj.Value)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid operation")
	}
	c.JSON(http.StatusNoContent, obj)
}

// @Summary delete key
// @Schemes http
// @Description delete key
// @Accept json
// @Produce json
// @Param RequestBody body dataDeleteSchema true "request body for delete an entry"
// @Success 200 {object} dataDeleteSchema
// @Router /core [delete]
func deleteHandler(c *gin.Context) {
	obj := dataDeleteSchema{}
	c.Bind(&obj)
	err := rebitcask.Delete(obj.Key)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid operation")
	}
	c.JSON(http.StatusAccepted, "")
}

func watchHandler(c *gin.Context) {}
