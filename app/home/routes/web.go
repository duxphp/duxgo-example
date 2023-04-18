package routes

import (
	"github.com/duxweb/go-fast/route"
	"github.com/gofiber/fiber/v2"
)

func RouteWeb(router *route.RouterData) {
	router.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("dsadsad")
	}, "首页", "home")
}
