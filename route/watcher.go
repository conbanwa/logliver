package route

import (
	"fmt"
	"log"

	"logliver/controller"
	"logliver/tool"

	"github.com/kataras/iris/v12"
)

func init() {
	AddRouter(func(app *iris.Application) {
		app.Get("/watcher", func(ctx iris.Context) {
			path := fmt.Sprintf(".%s", ctx.URLParam("pathname"))
			if !tool.IsExists(path) {
				log.Printf("Path not exists!")
				ctx.StatusCode(iris.StatusBadRequest)
				ctx.JSON("Path not exists!")
				return
			}
			ctx.Next()
		}, controller.WSWatcher.Serve())
	})
}
