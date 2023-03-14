package routes

import (
	"github.com/duxphp/duxgo/v2/route"
	"github.com/gofiber/fiber/v2"
)

func RouteWeb(router *route.RouterData) {
	router.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("dsadsad")
	}, "首页", "home")
}
