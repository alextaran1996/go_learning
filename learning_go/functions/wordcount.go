package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

var numofw int

func countwords(ref *html.Node) int {
	switch ref.Type {
	case html.ElementNode:
		for _, val := range ref.Attr {
			numofw += procstr(val.Val) // Add value when value field contains any value
			numofw += procstr(val.Key) // Add value when key field contains any value
		}
	case html.TextNode:
		numofw += procstr(ref.Data) // Add value from Data field

	}
	for child := ref.FirstChild; child != nil; child = child.NextSibling { // Recurcivly processing child elements
		countwords(child)
	}
	return numofw
	// Empty return. If we exactly specified returned values as (var_name type) we can use emty return

}

// Function that splits strings on words and count number of words
func procstr(s string) int {
	news := strings.Fields(s)
	return len(news)
}

func main() {
	resp, err := http.Get("https://golang.org") // input string | return *http.Response(All data retrived from get request)
	if err != nil {
		fmt.Println(err)
	}
	tree, err := html.Parse(resp.Body) //input io.Reader | return *html.Node(Pointer to root element of parsed html page)
	defer resp.Body.Close()            // Free up network resources
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(countwords(tree))

}
