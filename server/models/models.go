package models

import "gin-cmdb/server/middleware"

/**
 * @Description
 * @Author sjie
 * @Date 2022/6/17 17:58
 **/

var db = middleware.NewConnectMySQL()

func init() {
	// 初始化表
	err := db.AutoMigrate(&Users{}, &Host{})
	if err != nil {
		return
	}
}
