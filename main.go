package main

import (
	"github.com/prclin/minimal-tiktok/controller"
	"github.com/prclin/minimal-tiktok/core"
	"github.com/prclin/minimal-tiktok/global"
	"net/http"
	"strconv"
	"time"
)

func main() {
	//初始化项目核心
	core.Init()
	//加载路由
	controller.Init()

	//服务关闭时，将logger缓冲区中日志刷出
	defer global.Logger.Sync()

	//启动http服务器
	server := &http.Server{
		Addr:           ":" + strconv.Itoa(global.Configuration.Server.Port),
		Handler:        core.Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic("server start failed...")
	}
}
