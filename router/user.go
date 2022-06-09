package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @Description
 * @Author sjie
 * @Date 2022/6/9 14:16
 **/

func loadUserRouter(e *gin.Engine) {
	userApiGroup := e.Group("/user")
	{
		userApiGroup.POST("/login", loginUser)
	}
}

func loginUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	msg := fmt.Sprintf("username:%s, password:%s", username, password)
	fmt.Println(msg)
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  msg,
	})
}
