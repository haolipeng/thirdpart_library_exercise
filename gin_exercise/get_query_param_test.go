package gin_exercise

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

type Info struct {
	Name    string
	Age     string
	Company string
}

func handlerQueryParam(c *gin.Context) {
	//判断获取的值是否存在
	_, exist := c.GetQuery("male")
	if !exist {
		c.JSON(http.StatusInternalServerError, "key value is not exist")
		return
	}

	name := c.Query("name")
	age := c.Query("age")
	company := c.Query("company")

	c.JSON(http.StatusOK, Info{
		Name:    name,
		Age:     age,
		Company: company,
	})
	//也可以使用c.DefaultQuery()函数，在无法获取值时返回设置的默认值
}

//测试url为http://localhost:8080/users?name=haolipeng&age=31&company=antiy
func TestQueryParam(t *testing.T) {
	engine := gin.Default()
	engine.GET("/users", handlerQueryParam)

	engine.Run(":8080")
}
