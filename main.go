package main

import (
	"duxgotest/app/home"
	"embed"
	"github.com/duxphp/duxgo"
	admin "github.com/duxphp/duxgo-admin"
	ui "github.com/duxphp/duxgo-ui"
	"github.com/duxphp/duxgo/bootstrap"
	"github.com/duxphp/duxgo/core"
	"html/template"
)

//go:embed views/* app/*/views/*
var ViewsFs embed.FS

func main() {

	// 创建框架服务
	server := duxgo.New()

	// 设置基础配置
	server.SetConfigDir("./config/")
	server.SetDatabaseStatus(true)
	server.SetRedisStatus(true)
	server.SetQueueStatus(true)
	server.SetWebsocketStatus(true)
	server.SetMongodbStatus(false)

	// 注册应用模块
	server.RegisterApp(home.App, admin.New)

	// 注册框架服务
	server.RegisterService(ui.New, func(bootstrap *bootstrap.Bootstrap) {
		template.Must(core.Tpl.ParseFS(ViewsFs, "views/*", "app/*/views/*"))
	})

	// 启动框架服务
	server.Start()
}
