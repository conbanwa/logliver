package main

import (
	"fmt"
	"io/ioutil"
	"logliver/route"
	"net/http"

	"github.com/JabinGP/demo-chatroom/middleware"
	"github.com/gin-gonic/gin"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	// router1 := gin.Default()
	// router1.Use(CrosHandler())
	// // router1.Use(static.Serve("/", static.LocalFile("web", false)))
	// router1.GET("/path", Directory)
	// router1.Any("/watcher", func(ctx iris.Context) {
	// 	path := fmt.Sprintf(".%s", ctx.URLParam("pathname"))
	// 	if !tool.IsExists(path) {
	// 		log.Printf("Path not exists!")
	// 		ctx.StatusCode(iris.StatusBadRequest)
	// 		ctx.JSON("Path not exists!")
	// 		return
	// 	}
	// 	ctx.Next()
	// }, Serve)

	// router1.Run(":8080")
	app := iris.New()

	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	app.AllowMethods(iris.MethodOptions)
	app.Use(middleware.CORS)

	route.Route(app)

	app.Run(iris.Addr("localhost:1234"), iris.WithoutServerError(iris.ErrServerClosed))
}

func CrosHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(204)
			return
		}
		context.Next()
	}
}

func Directory(c *gin.Context) {
	path := c.Query("path")
	files, err := ioutil.ReadDir("/logs/" + path)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, files)
}

// func Serve(c *gin.Context) {
// 	ch := channelAdd(c.Query("id") + c.Query("version"))

// 	wsHandler := func(ws *websocket.Conn) {

// 		go func() {

// 		}()

// 		if err != nil {
// 			fmt.Println(err.Error())
// 			return
// 		}

// 		<-ch
// 		wrapper.Close()
// 		ch <- false
// 	}
// 	defer log.Printf("Websocket session closed for %v", c.Request.RemoteAddr)
// 	websocket.Handler(wsHandler).ServeHTTP(c.Writer, c.Request)
// }
