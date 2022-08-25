package home

import (
	"duxgotest/app/home/routes"
	"github.com/duxphp/duxgo/register"
	"github.com/duxphp/duxgo/util"
	"github.com/labstack/echo/v4"
)

var config = struct {
}{}

func App() {
	register.App(&register.AppConfig{
		Name:     "home",
		Config:   &config,
		Route:    Route,
		AppRoute: AppRoute,
	})
}

func AppRoute(router register.Router, app *echo.Echo) {
	router["web"] = util.NewRouter(app.Group(""))

}

func Route(router register.Router) {
	routes.RouteWeb(router["web"])
}
