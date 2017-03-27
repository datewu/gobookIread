package main

import (
	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/view"
)

func main() {
	// Configuration is optional
	app := iris.New(iris.Configuration{Gzip: false, Charset: "UTF-8"})

	// Adapt a logger which will print all errors to os.Stdout
	app.Adapt(iris.DevLogger())

	// Adapt the httprouter (we will use that on all examples
	app.Adapt(httprouter.New())

	// Parse all files inside `./mytemplates` directory ending with `.html`
	app.Adapt(view.HTML("./mytemplates", ".html"))

	todos := []Todo{
		{"Learn Go", true},
		{"Read GopuerBook", true},
		{"Create a web app in Go", false},
	}

	app.Get("/", func(ctx *iris.Context) {
		//ctx.Render("todos.html", struct{ Data []Todo }{todos})
		ctx.Render("todos.html", todos)
	})

	app.Listen(":8080")

}

// Todo bind struct
type Todo struct {
	Task string
	Done bool
}
