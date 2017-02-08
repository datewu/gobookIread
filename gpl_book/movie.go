package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	//data, err := json.Marshal(movies)
	data, err := json.MarshalIndent(movies, "", "   ")
	if err != nil {
		log.Println("JSON marshaling failed:", err)
		return
	}
	fmt.Printf("%s\n", data)

	var titles []struct{ Title string }
	if err = json.Unmarshal(data, &titles); err != nil {
		log.Println("JSON unmarshaling failed:", err)
		return
	}

	fmt.Println(titles)
}

type movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}
