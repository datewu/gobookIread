package main

import (
	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	app := iris.New()
	app.Adapt(httprouter.New())

	// first parameter is the request path
	// second is the operationg system directory
	app.StaticWeb("/static", "./assets")
	app.Listen(":8081")
}
