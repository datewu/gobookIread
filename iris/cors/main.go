package main

import (
	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/cors"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	app := iris.New()
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	app.Adapt(crs) // this line should be added
	// adaptor supports cors allowed methods, middleware does not.

	// if you want pre-route-only cors
	// check middleware  instead

	v1 := app.Party("/api/v1")
	{
		v1.Post("/home", func(c *iris.Context) {
			app.Log(iris.DevMode, "lol")
			c.WriteString("response Post from /home")
		})
		v1.Get("/g", func(c *iris.Context) {
			app.Log(iris.DevMode, "cf")
			c.WriteString("response get from /g")
		})
		v1.Post("/h", func(c *iris.Context) {
			app.Log(iris.DevMode, "dota")
			c.WriteString("response get from /h")
		})
	}
	app.Listen(":8080")

}
