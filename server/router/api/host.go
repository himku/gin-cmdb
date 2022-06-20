package api

import (
	"gin-cmdb/server/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
 * @Description
 * @Author sjie
 * @Date 2022/6/20 10:14
 **/

func LoadHostRouter(e *gin.Engine) {
	hostApiGroup := e.Group("/host")
	{
		hostApiGroup.GET("/list", ListHost)
	}
}

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
