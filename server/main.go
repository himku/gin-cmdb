package main

import (
	"fmt"
	"gin-cmdb/server/config"
	"gin-cmdb/server/router"
	"net/http"
	"time"
)

func main() {
	initRouter := router.InitRouter()
	initConfig := config.NewConfig()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", initConfig.Server.Port),
		Handler:        initRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		return
	}
}
