package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @Description
 * @Author sjie
 * @Date 2022/6/6 16:19
 **/

type Auth struct {
	Username string `gorm:"username"`
	Password string `gorm:"password"`
}

func CheckAuth(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	msg := fmt.Sprintf("username:%s, password:%s,", username, password)
	Cat := fmt.Sprintf("hello:%s", "s")
	fmt.Println(msg)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  msg,
		"cat":  Cat,
	})
}
