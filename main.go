package main

import (
	"logliver/route"

	"github.com/JabinGP/demo-chatroom/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	app.AllowMethods(iris.MethodOptions)
	app.Use(middleware.CORS)

	route.Route(app)

	app.Run(iris.Addr("10.238.4.33:1234"), iris.WithoutServerError(iris.ErrServerClosed))
}
