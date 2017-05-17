package main

import (
	"time"

	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/middleware/basicauth"
)

func main() {
	app := iris.New()
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())

	authConfig := basicauth.Config{
		Users:      map[string]string{"lol": "good", "dota": "nice"},
		Realm:      "Authorization Required, due",
		ContextKey: "mycustomkey",
		Expires:    time.Duration(30) * time.Minute,
	}

	authentication := basicauth.New(authConfig)
	app.Get("/", func(ctx *iris.Context) {
		ctx.Redirect("/admin")
	})

	// to global
	// app.Use(authentication)

	// to routes
	/*
		app.Get("/mysecret", authentication, func(ctx *iris.Context) {
			username := ctx.GetString("mycustomkey")
			ctx.Writef("Hello authenticated user: %s ", username)
		})
	*/

	// to party
	needAuth := app.Party("/admin", authentication)
	{
		// /admin
		needAuth.Get("/", func(ctx *iris.Context) {
			username := ctx.GetString("mycustomkey")
			ctx.Writef("Hello authenticated user: %s ", username)

		})

		// /admin/profile
		needAuth.Get("/profile", func(ctx *iris.Context) {
			username := ctx.GetString("mycustomkey")
			ctx.Writef("Hello authenticated user: %s ", username)

		})

		// /admin/settings
		needAuth.Get("/settings", func(ctx *iris.Context) {
			username := ctx.GetString("mycustomkey")
			ctx.Writef("Hello authenticated user: %s ", username)

		})
	}
	app.Listen(":7070")
}
