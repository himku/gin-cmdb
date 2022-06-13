package router

import (
	"fmt"
	"gin-cmdb/models"
	"gin-cmdb/utils"
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
		userApiGroup.POST("/login", LoginUser)
		userApiGroup.POST("/add", CreateUser)
	}
}

func LoginUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if models.CheckAuth(username, password) {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"msg":    "登陆成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"msg":    "登陆失败",
		})
	}

}

func CreateUser(c *gin.Context) {
	//  获取表单数据
	username := c.PostForm("username")
	password := c.PostForm("password")
	hashPassword, _ := utils.HashPassword(password)
	user := models.Users{Username: username, Password: hashPassword}
	if models.CreateUser(&user) {
		msg := fmt.Sprintf("用户%s创建成功", username)
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"msg":    msg,
		})
	} else {
		msg := fmt.Sprintf("用户%s已存在!", username)
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"msg":    msg,
		})
	}
}
