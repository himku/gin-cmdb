package main

import (
	"fmt"
	"gin-cmdb/server/config"
	"gin-cmdb/server/router"
)

func main() {
	r := router.InitRouter()
	c := config.NewConfig()
	serverPort := fmt.Sprintf(":%s", c.Server.Port)
	err := r.Run(serverPort)
	if err != nil {
		fmt.Printf("服务启动失败:%s", err)
	}
}
