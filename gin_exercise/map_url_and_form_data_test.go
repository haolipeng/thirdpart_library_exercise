package gin_exercise

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

/*
测试url 推荐使用postman
POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
Content-Type: application/x-www-form-urlencoded

names[first]=thinkerou&names[second]=tianou
*/

func handleMapUrlParamAndFormData(c *gin.Context) {
	ids := c.QueryMap("ids")
	names := c.PostFormMap("names")
	fmt.Printf("ids:%v,names:%v\n", ids, names)
}

func TestMapUrlParamAndFormData(t *testing.T) {
	engine := gin.Default()

	engine.POST("/post", handleMapUrlParamAndFormData)

	engine.Run(":8080")
}
