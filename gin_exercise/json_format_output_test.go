package gin_exercise

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
}

//普通json格式输出
func TestJsonFormatOutput(t *testing.T) {
	r := gin.Default()
	s1 := Student{
		Name: "haolipeng",
		Age:  32,
	}
	s2 := Student{
		Name: "zhouyang",
		Age:  33,
	}
	var stuList []Student
	stuList = append(stuList, s1, s2)

	r.GET("/users/haolipeng", func(c *gin.Context) {
		c.JSON(http.StatusOK, &Response{
			Data:    stuList,
			Message: "success",
			Status:  http.StatusOK,
		})
	})
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("server run failed")
	}
}

type BusinessInfo struct {
	IpAddress        string `json:"ipAddress"`        //ip地址
	Hostname         string `json:"hostname"`         //主机名称
	SystemName       string `json:"systemName"`       //操作系统
	HostStatus       string `json:"hostStatus"`       //在线状态
	Grade            string `json:"grade"`            //资产等级
	GroupName        string `json:"groupName"`        //业务组
	HostType         string `json:"type"`             //主机类型（lin为Linux、win为Windows）
	LabelBusiness    string `json:"labelBusiness"`    //业务标签
	LabelPosition    string `json:"labelPosition"`    //位置标签
	LabelEnvironment string `json:"labelEnvironment"` //环境标签
	LabelRole        string `json:"labelRole"`        //角色标签
}

var assetsInfos = []BusinessInfo{
	{
		IpAddress:        "192.168.1.1",
		Hostname:         "OAWeb",
		SystemName:       "CentOS Linux",
		HostStatus:       "离线",
		Grade:            "重要资产",
		HostType:         "lin",
		GroupName:        "北京组",
		LabelPosition:    "北京",
		LabelBusiness:    "ERP系统",
		LabelEnvironment: "生产",
		LabelRole:        "Nginx",
	},
	{
		IpAddress:        "192.168.1.2",
		Hostname:         "OAWeb",
		SystemName:       "CentOS Linux",
		HostStatus:       "离线",
		Grade:            "重要资产",
		HostType:         "lin",
		GroupName:        "北京组",
		LabelPosition:    "北京",
		LabelBusiness:    "ERP系统",
		LabelEnvironment: "研发",
		LabelRole:        "Nginx",
	},
	{
		IpAddress:        "192.168.1.3",
		Hostname:         "OAWeb",
		SystemName:       "CentOS Linux",
		HostStatus:       "离线",
		Grade:            "重要资产",
		HostType:         "lin",
		GroupName:        "北京组",
		LabelPosition:    "北京",
		LabelBusiness:    "ERP系统",
		LabelEnvironment: "销售",
		LabelRole:        "Nginx",
	},
	{
		IpAddress:        "192.168.1.4",
		Hostname:         "OAWeb",
		SystemName:       "CentOS Linux",
		HostStatus:       "离线",
		Grade:            "重要资产",
		HostType:         "lin",
		GroupName:        "北京组",
		LabelPosition:    "北京",
		LabelBusiness:    "ERP系统",
		LabelEnvironment: "财务",
		LabelRole:        "Nginx",
	},
}

func handleAssetsJson(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, assetsInfos)
}

func handleTask(c *gin.Context) {
	resp := Response{
		Data:    "1004",
		Message: "success",
		Status:  0,
	}
	c.JSON(http.StatusOK, resp)
}

func TestJsonArrayOutput(t *testing.T) {
	r := gin.Default()
	r.POST("/base/host/get/attribute", handleAssetsJson)
	err := r.Run(":8084")
	if err != nil {
		fmt.Println("server run failed")
	}
}

func TestMockAgentProxyService(t *testing.T) {
	r := gin.Default()
	r.POST("/console/task/create_new", handleTask)
	err := r.Run(":8084")
	if err != nil {
		fmt.Println("server run failed")
	}
}
