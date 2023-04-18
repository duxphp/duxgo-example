package home

import (
	"github.com/duxweb/go-fast/app"
	"github.com/duxweb/go-fast/route"
	"github.com/gofiber/fiber/v2"
	"github.com/rotisserie/eris"
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

	group.Get("/test", func(c *fiber.Ctx) error {
		return eris.New("dsad")
		return c.SendString("I'm a GET request!")
	}, "首页", "web.home")

}
