package gin_exercise

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func handleParamArray(c *gin.Context) {
	arr := c.QueryArray("media")
	c.JSON(http.StatusOK, arr)
}

//测试url为http://localhost:8080/?media=blog&media=wechat
func TestQueryArray(t *testing.T) {
	engine := gin.Default()
	engine.GET("/", handleParamArray)
	engine.Run(":8080")
}
