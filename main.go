package main

import (
	"duxgotest/app/home"
	"github.com/duxphp/duxgo"
	ui "github.com/duxphp/duxgo-ui"
)

func main() {
	server := duxgo.New()
	server.SetConfigDir("./config/")
	server.SetDatabaseStatus(false)
	server.SetRedisStatus(false)
	server.SetQueueStatus(false)
	server.SetWebsocketStatus(false)
	server.SetMongodbStatus(false)
	server.RegisterApp(home.App)
	server.RegisterService(ui.New)
	server.Start()
}
