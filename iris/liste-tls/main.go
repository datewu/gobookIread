package main

import (
	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

const host = "127.0.0.1:54443"

func main() {
	app := iris.New()
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())

	app.Get("/", func(ctx *iris.Context) {
		ctx.Writef("Hello from the SECURE server")
	})

	app.Get("/mypath", func(ctx *iris.Context) {
		ctx.Writef("Hello from the SECURE server on /mypath")
	})

	// start a secondary server (http) on port 8080, this is a non-blocking func
	// redirect all http to the main server which is tls/ssl on port :443
	iris.Proxy(":8080", "https://"+host)

	// Start the MAIN server (HTTPS) on host, this is a blocking func
	app.ListenTLS(host, "server.cert", "server.key")
}
