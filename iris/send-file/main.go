package main

import (
	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	app := iris.New()
	// output startip banner and err logs on os.Stdout
	app.Adapt(iris.DevLogger())

	// set the router, you can choose gorillamux too
	app.Adapt(httprouter.New())

	app.Get("/dl", func(c *iris.Context) {
		file := "./files/lol.txt"
		c.SendFile(file, "dota.html")
	})
	app.Listen(":8080")
}
