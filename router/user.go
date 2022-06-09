package router

import (
	"GoNotes/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Description
 * @Author sjie
 * @Date 2022/6/9 14:16
 **/

func loadUserRouter(e *gin.Engine) {
	// 路由分组
	userApiGroup := e.Group("/user")
	{
		userApiGroup.POST("/add", CreateUser)
	}
}

func CreateUser(c *gin.Context) {
	//  获取表单数据
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := models.Users{Username: username, Password: password}
	models.CreateUser(&user)
	models.CreateUser(&user)
	if models.CreateUser(&user) {
		msg := fmt.Sprintf("create username %s failed!", username)
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"msg":    msg,
		})
	} else {
		msg := fmt.Sprintf("create username %s success!", username)
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"msg":    msg,
		})
	}
}
