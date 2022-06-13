package middleware

import (
	"gin-cmdb/server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

/**
 * @Description
 * @Author sjie
 * @Date 2022/6/13 15:37
 **/

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "未携带认证Token",
			})
			c.Abort() // 结束后续请求操作
			return
		}
		headerParts := strings.SplitN(authHeader, " ", 2)
		if !(len(headerParts) == 2 && headerParts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "header Auth 格式错误",
			})
			c.Abort()
			return
		}
		claims, err := utils.ParseClamsToken(headerParts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "无效Token",
			})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next() // 交给后面处理
	}
}
