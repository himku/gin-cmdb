package router

import (
	"gin-cmdb/server/router/user"
	"github.com/gin-gonic/gin"
)

/**
 * @Description
 * @Author sjie
 * @Date 2022/6/7 17:10
 **/

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	user.LoadUserRouter(r)
	return r
}
