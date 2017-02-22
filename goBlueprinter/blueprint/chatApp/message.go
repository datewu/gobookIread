package main

import "time"

// message repressents a single meessage
type message struct {
	Name      string
	Message   string
	When      time.Time
	AvatarURL string
}
