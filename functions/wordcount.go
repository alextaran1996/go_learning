package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func countwords(ref *html.Node) (numofw int, err error) {
	if ref.Type == html.ElementNode {
		for _, val := range ref.Attr {
			numofw += procstr(val.Val)
		}
		for child := ref.FirstChild; child != nil; child = child.NextSibling {
			countwords(child)
		}
	}
	return numofw, nil
}

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
