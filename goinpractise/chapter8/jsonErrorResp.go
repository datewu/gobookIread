package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Error lol
type Error struct {
	HTTPCode int `json:"-"`
	Code     int `json:"code,omitempty"`
	Message  string
}

// JSONError lol
func JSONError(w http.ResponseWriter, e Error) {
	data := struct {
		Err Error `json:"error"`
	}{e}
	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.HTTPCode)
	fmt.Fprint(w, string(b))
}

func displayError(w http.ResponseWriter, r *http.Request) {
	e := Error{
		HTTPCode: http.StatusForbidden,
		Code:     344,
		Message:  "An error Occurred, opps",
	}
	JSONError(w, e)
}

func main() {
	http.HandleFunc("/", displayError)
	http.ListenAndServe(":8080", nil)
}
