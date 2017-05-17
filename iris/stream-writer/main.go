package main

import (
	"fmt"
	"io"
	"time"

	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	app := iris.New()
	// output start banner and error logs on os.Stdout
	app.Adapt(iris.DevLogger())
	// set the router, you can choose gorillamux too
	app.Adapt(httprouter.New())

	timeWaitForCloseStream := 5 * time.Second

	app.Get("/", func(ctx *iris.Context) {
		i := 0
		// goroutine in order to no block and just wait,
		// goroutine is OPTIONAL and not a very good option but it depend
		// on the needs
		// look the streaming_simple_2 for an alternative code style
		// Send the response in chinks and wait for a second between
		// each chunk.
		go ctx.StreamWriter(func(w io.Writer) bool {
			i++
			fmt.Fprintf(w, "this is a message number %d\n", i) // write
			time.Sleep(500 * time.Millisecond)
			if i == 8 {
				return false // close and flush
			}
			return true // continue write
		})

		// when this handler finished the client should be see the stream
		// writer's contens
		// simulate a job here..
		time.Sleep(timeWaitForCloseStream)
	})

	app.Get("/alternative", func(ctx *iris.Context) {
		// send the response in chunks and wait for a second bewteen each
		// chunk.
		ctx.StreamWriter(func(w io.Writer) bool {
			for i := 1; i <= 4; i++ {
				fmt.Fprintf(w, "this is a message number %d\n", i) //write
				time.Sleep(time.Second)
			}
			// when this handler finished the client should be see
			// the stream writer's contents
			return false // stop and flush the contents

		})

	})

	app.Listen(":8080")

}
