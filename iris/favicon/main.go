package main

import (
	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	app := iris.New()
	app.Adapt(httprouter.New())

	// This will server the ./static/fa.ico to localhot:8080/favicon.ico
	app.Favicon("./static/fa.ico")

	/*
		app.Favicon("./static/fa.ico", "/fa.ico") will
		server ./static/fa.icon to localhost:/fa.ico
	*/

	app.Get("/", func(ctx *iris.Context) {
		ctx.HTML(iris.StatusOK, `You should see the favicon now at the
		side of your browser, if not, please refresh or clear the browser
		cache`)
	})

	app.Listen(":8080")
}
