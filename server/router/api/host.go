package api

import (
	"fmt"
	"gin-cmdb/server/middleware"
	"gin-cmdb/server/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// LoadHostRouter /**
// 加载主机api router
func LoadHostRouter(e *gin.Engine) {
	hostApiGroup := e.Group("/host")
	{
		hostApiGroup.Use(middleware.JwtAuth())
		hostApiGroup.GET("/list", ListHost)
		hostApiGroup.POST("/add", AddHost)
		hostApiGroup.DELETE("/delete", DeleteHost)
		hostApiGroup.POST("/edit", ModifyHost)
	}
}

func ModifyHost(c *gin.Context) {
	hostName := c.PostForm("hostname")
	os := c.PostForm("os")
	cpuCores, _ := strconv.Atoi(c.PostForm("cpuCores"))
	status, _ := strconv.ParseBool(c.PostForm("status"))
	modifyHost := models.Host{OS: os, Hostname: hostName, CpuCores: cpuCores, Status: status}
	if models.EditHost(&modifyHost) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  fmt.Sprintf("主机%s修改成功", hostName),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  fmt.Sprintf("主机%s修改失败", hostName),
		})
	}
}

func DeleteHost(c *gin.Context) {
	hostName := c.Query("hostname")
	if models.DeleteHost(hostName) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  fmt.Sprintf("主机%s删除成功", hostName),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  fmt.Sprintf("主机%s删除失败", hostName),
		})
	}
}

// ListHost 主机列表路由函数
func ListHost(c *gin.Context) {
	hostPageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
	page, _ := strconv.Atoi(c.PostForm("page"))
	total, hostList := models.GetHostList(page, hostPageSize)
	if (total + len(hostList)) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "未获取到主机数据",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  hostList,
		})
	}
}

// 	AddHost 添加主机路由函数
func AddHost(c *gin.Context) {
	hostName := c.PostForm("hostname")
	os := c.PostForm("os")
	cpuCores, _ := strconv.Atoi(c.PostForm("cpuCores"))
	status, _ := strconv.ParseBool(c.PostForm("status"))
	host := models.Host{OS: os, Hostname: hostName, CpuCores: cpuCores, Status: status}
	if models.CreateHost(&host) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  fmt.Sprintf("host %s 创建成功", host.Hostname),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  fmt.Sprintf("host %s 已经存在", host.Hostname),
		})
	}

}
