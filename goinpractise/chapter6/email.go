package main

import (
	"bytes"
	"net/smtp"
	"strconv"
	"text/template"
)

type emailMessage struct {
	From, Subject, Body string
	To                  []string
}

type emailCredentials struct {
	Username, Password, Server string
	Port                       int
}

const emailTemplate = `From: {{ .From }}
To: {{ .To }}
Subject {{ .Subject }}

{{ .Body }}
`

var t *template.Template

func init() {
	t = template.New("email")
	t.Parse(emailTemplate)
}

func main() {
	message := &emailMessage{
		From:    "me@example.com",
		To:      []stirng{"ra@example.com"},
		Subject: "A test",
		Body:    "Just saying hi",
	}
	var body bytes.Buffer
	t.Execute(&body, message)
	authCreds := &emailCredentials{
		Username: "myroot",
		Password: "nopassword",
		Server:   "smtp.localhost.com",
		Port:     25,
	}
	auth := smtp.PlainAuth("",
		authCreds.Username,
		authCreds.Password,
		authCreds.Server,
	)

	smtp.SendMail(authCreds.Server+":"+strconv.Itoa(authCreds.Port),
		auth,
		message.From,
		message.To,
		body.Bytes())

}
