package router

import (
	"fmt"
	"gin-cmdb/server/middleware"
	"gin-cmdb/server/models"
	"gin-cmdb/server/utils"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
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
		userApiGroup.Use(middleware.JwtAuth())
		userApiGroup.POST("/refreshToken", RefreshAuth)
		userApiGroup.POST("/add", CreateUser)
		userApiGroup.GET("/list", ListUser)
	}
}

func RefreshAuth(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	refreshToken := strings.SplitN(authHeader, " ", 2)
	newToken, _ := utils.RefreshToken(refreshToken[1])
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  newToken,
	})
}

func ListUser(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
	page, _ := strconv.Atoi(c.PostForm("page"))
	total, userList := models.GetUserList(page, pageSize)
	if (total + len(userList)) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "未获取到用户数据",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  userList,
	})
}

func LoginUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if models.CheckAuth(username, password) {
		user := utils.JwtCustomClaims{Username: username, Password: password, StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Minute * 60)),
		}}
		token, err := utils.MakeClamsToken(user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"msg":    "登陆成功",
			"data":   gin.H{"token": token},
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
