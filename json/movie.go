package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`        // If value not specified  output released: 0
	Color  bool `json:"color,omitempty"` // If value not specified it'll not appear in output
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
		{Title: "Bullit", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
		{Title: "Avatar", Actors: []string{"Joel David Moore", "Giovanni Ribisi"}},
	}
	data, err := json.MarshalIndent(movies, "", " ")
	if err != nil {
		fmt.Println("Problems while processing data")
	}
	fmt.Printf("%s\n", data)
}
