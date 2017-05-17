package main

import (
	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/middleware/logger"
)

func main() {
	app := iris.New()

	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())

	customLogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
	})

	app.Use(customLogger)

	app.Get("/", func(ctx *iris.Context) {
		ctx.Writef("hello homepage")
	})

	app.Get("/1", func(ctx *iris.Context) {
		ctx.Writef("hello path 1")
	})

	app.Get("/2", func(ctx *iris.Context) {
		ctx.Writef("hello path 2")
	})

	// log http errors
	errorLogger := logger.New()

	app.OnError(404, func(ctx *iris.Context) {
		errorLogger.Serve(ctx)
		ctx.Writef("My custom 404 error page")
	})

	app.Listen(":8085")

}
