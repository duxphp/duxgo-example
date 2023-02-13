package home

import (
	"github.com/duxphp/duxgo/v2/app"
)

var config = struct {
}{}

func App() {
	app.Register(&app.Config{
		Name:   "home",
		Title:  "演示",
		Desc:   "这是一个演示应用",
		Config: &config,
		Init:   Init,
	})
}

func Init() {
}
