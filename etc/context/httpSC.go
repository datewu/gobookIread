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
	http.HandleFunc("/value", contextValue(value))
	http.HandleFunc("/tout", contextDead(line))
	http.HandleFunc("/", hw)
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
			fmt.Fprintf(w, "timeout\n")
			break lo
		}

	}
	fmt.Fprintf(w, " outage")
}

type k string

const key k = "key dota"

func contextValue(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), key, "lol")
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func contextDead(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 500*time.Millisecond)
		defer cancel()
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
