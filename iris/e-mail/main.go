package main

import (
	"bytes"
	"html/template"
	"path/filepath"
	"strconv"

	mailer "github.com/kataras/go-mailer"
	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

var t = template.Must(template.ParseFiles(filepath.Join("templates", "body.html")))

func main() {
	app := iris.New()
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())

	// change those to your own settings
	cfg := mailer.Config{
		Host:     "smtp.mailgun.org",
		Username: "postmaster@sandbox661c307650f04e909150b37c0f3b2f09.mailgun.org",
		Password: "38304272b8ee5c176d5961dc155b2417",
		Port:     587,
	}

	// create the service
	mailService := mailer.New(cfg)

	var to = []string{"hihahajun@gmail.com"}

	// standalone
	// mailService.Send("iris email test subject", "<h1> outside of
	// context before server's listen!</h1>", to...)

	// insider handler
	app.Get("/send", func(ctx *iris.Context) {
		content := `<h1> Hello From Iris Web Framework</h1>
		<br/><br/> <span style="color:blue"> This is the rich message 
		body</span>`
		err := mailService.Send("iris e-mail just test subject", content, to...)

		if err != nil {
			ctx.HTML(503, "<b> Problem while sending the email: "+err.Error())
		} else {
			ctx.HTML(200, "<h1> SUCCESS! Checkout you inbox :> </h1>")
		}
	})

	// send a body by template
	app.Get("/send/template", func(ctx *iris.Context) {
		buf := &bytes.Buffer{}

		data := map[string]string{
			"Message": "This is the rich message again",
			"Footer":  "The footer of the email!",
		}
		t.Execute(buf, data)

		content := buf.String()

		ctx.HTML(200, content+"\n LOL \n"+strconv.Itoa(len(content)))
		/*
			err := mailService.Send("iris e-mail just test subject", content, to...)

			if err != nil {
				ctx.HTML(503, "<b> Problem while sending the email: "+err.Error())
			} else {
				ctx.HTML(200, "<h1> SUCCESS! Checkout you inbox :> </h1>")
			}
		*/
	})

	app.Listen(":8080")
}
