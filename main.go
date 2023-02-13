package main

import (
	"duxgotest/app/home"
	"embed"
	"github.com/duxphp/duxgo/v2"
)

//go:embed views/* app/*/views/*
var ViewsFs embed.FS

func main() {

	// 创建框架服务
	dux := duxgo.New()

	dux.RegisterApp(home.App)

	dux.Run()

	//// 设置基础配置
	//server.SetConfigDir("./config/")
	//server.SetDatabaseStatus(true)
	//server.SetRedisStatus(true)
	//server.SetQueueStatus(true)
	//server.SetWebsocketStatus(true)
	//server.SetMongodbStatus(false)
	//
	//// 注册应用模块
	//server.RegisterApp(home.App, admin.New)
	//
	//// 注册框架服务
	//server.RegisterService(ui.New, func(bootstrap *bootstrap.Bootstrap) {
	//	template.Must(core.Tpl.ParseFS(ViewsFs, "views/*", "app/*/views /*"))
	//})
	//
	//// 启动框架服务
	//server.Start()
}
