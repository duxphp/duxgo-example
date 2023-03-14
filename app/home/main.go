package home

import (
	"fmt"
	"github.com/duxphp/duxgo/v2/app"
	"github.com/duxphp/duxgo/v2/global"
	"github.com/duxphp/duxgo/v2/route"
	"github.com/gofiber/fiber/v2"
)

var config = struct {
}{}

func App() {
	app.Register(&app.Config{
		Name:     "home",
		Title:    "演示",
		Desc:     "这是一个演示应用",
		Config:   &config,
		Init:     Init,
		Register: Register,
	})
}

func Init() {
	route.Add("web", route.New(""))
}

func Register() {
	group := route.Get("web")
	group.Get("/ss", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	}, "首页", "web.home")

	routes := global.App.GetRoutes()
	for _, r := range routes {
		fmt.Printf("%s - %s\n", r.Method, r.Path)
	}

}
