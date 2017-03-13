package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	//http.HandleFunc("/", hw)
	http.Handle("/value", contextValue(value))
	http.Handle("/tout", contextDead(line))
	http.Handle("/", http.HandlerFunc(hw))
	http.Handle("/", hw)
	log.Println(http.ListenAndServe(":8080", nil))
}

func hw(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func value(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.Context().Value(key).(string))
}

func line(w http.ResponseWriter, r *http.Request) {
lo:
	for {
		select {
		case <-time.After(5 * time.Second):
			fmt.Fprintf(w, "hello 5s")
			break lo
		case <-r.Context().Done():
			fmt.Fprintf(w, "timeout")
			break lo
		}

	}
	fmt.Fprintf(w, " outage")
}

type k string

const key k = "key dota"

func contextValue(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), key, "lol")
			next.ServeHTTP(w, r.WithContext(ctx))
		})
}

func contextDead(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), 250*time.Millisecond)
			defer cancel()
			next.ServeHTTP(w, r.WithContext(ctx))
		})
}
