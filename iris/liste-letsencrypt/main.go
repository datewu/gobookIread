package main

import (
	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	app := iris.New()
	app.Adapt(iris.DevLogger)
	app.Adapt(httprouter.New())

	app.Get("/", func(ctx *iris.Context) {
		ctx.Writef("hello from SECURE SERVER!")
	})

	app.Get("/test2", func(ctx *iris.Context) {
		ctx.Writef("Welcome to secure server from /test2!")
	})

	app.Get("/redirect", func(ctx *iris.Context) {
		ctx.Redirect("/test2")
	})

	app.ListenLETSENCRYPT("wutuofu.com:443")
	// This will provide you automatic certification & key from
	// letsencrypt.org's servers
	// it also starts a second http:// server which redirect all
	// http://$PATH requests to https://$PATH
}
