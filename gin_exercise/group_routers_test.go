package gin_exercise

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func handleV1(c *gin.Context) {
	//所有/v1/xx的请求都会走这个地方，方便/v1分组路由下的授权校验处理
	fmt.Println("统一处理v1请求")
}

//温馨提示，一般可以用{}把不同分组的括起来
func TestGroupRoutes(t *testing.T) {
	r := gin.Default()
	v1Group := r.Group("/v1", handleV1)
	v1Group.GET("/users", func(c *gin.Context) {
		c.String(http.StatusOK, "v1/users")
	})

	v1Group.GET("/products", func(c *gin.Context) {
		c.String(http.StatusOK, "v1/products")
	})

	v2roup := r.Group("/v2")
	v2roup.GET("/users", func(c *gin.Context) {
		c.String(http.StatusOK, "v2/users")
	})

	v2roup.GET("/products", func(c *gin.Context) {
		c.String(http.StatusOK, "v2/products")
	})

	r.Run(":8080")
}
