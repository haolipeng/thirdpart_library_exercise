package gin_exercise

import (
	"github.com/gin-gonic/gin"
	"testing"
)

type User struct {
	Age  int
	Name string
}

var userList = []User{{Name: "haolipeng", Age: 31}, {Name: "zhouyang", Age: 32}}

func handlerUrlParam(c *gin.Context) {
	c.JSON(200, userList)
}

func handleUserId(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "user id:%s", id)
}

//测试url为http://localhost:8080/user/:haolipeng
func TestUrlParam(t *testing.T) {
	engine := gin.Default()
	engine.GET("/users", handlerUrlParam)
	engine.GET("/users/:id", handleUserId)
	engine.Run(":8080")
}
