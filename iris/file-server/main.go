package main

import (
	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	app := iris.New()
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())

	// first parameter is the request path
	// second is the operating system directory
	app.StaticWeb("/s", "./assets")
	app.Listen(":8080")
}
