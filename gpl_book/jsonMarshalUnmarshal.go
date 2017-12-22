package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var movies = []movie{
		{"Casablance", 1942, false,
			[]string{"Humphrey Bogart", "Ingrid Bergman"},
		},
		{"Cool Hand Luke", 1967, true,
			[]string{"Paul Newman"},
		},
		{"Bullitt", 1968, true,
			[]string{"Steve McQueen", "Jacqueline Bisset"},
		},
	}

	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalln("JSON Marshaling failed", err)
	}
	fmt.Printf("%s\n", data)

	var titles []struct{ Title string }
	if err = json.Unmarshal(data, &titles); err != nil {
		log.Fatalln("JSON unmarshling failed", err)
	}
	fmt.Println(titles)
}

type movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}
