package main

import (
	"log"

	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/view"
)

// ContactDetails the information from user
type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	app := iris.New()
	app.Adapt(httprouter.New())

	// Parse all files inside `./tem` directory ending wiht `.html`
	app.Adapt(view.HTML("./tem", ".html"))

	app.Get("/", func(ctx *iris.Context) {
		ctx.Render("forms.html", nil)
	})

	// Equivalent with app.HandleFunc("POST", ...)
	app.Post("/", func(ctx *iris.Context) {
		/*
			details := ContactDetails{
				Email:   ctx.FormValue("Email"),
				Subject: ctx.FormValue("Subject"),
				Message: ctx.FormValue("Message"),
			}
		*/
		// or simple:
		var details ContactDetails
		ctx.ReadForm(&details)

		ctx.Render("forms.html", details)
		log.Println(details, "and", details.Email)
	})
	app.Listen(":8080")
}
