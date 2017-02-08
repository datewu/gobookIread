package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 99, "sock": 98}
	log.Fatalln(http.ListenAndServe(":8080", db))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
