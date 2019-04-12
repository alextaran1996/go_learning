package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// Get all links founded on the page
func readlinks(links []string, n *html.Node) []string {

	for _, val := range n.Attr {
		if val.Key == "href" {
			links = append(links, val.Val)
		}
	}
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		links = readlinks(links, child)
	}
	return links
}

const htmlURL = "https://golang.org"

func main() {
	resp, err := http.Get(htmlURL) // Return *http.Response type of reader Reader.Included all the fields from Response structure
	if err != nil {
		fmt.Println("Can not open URL: ", err)
	}
	// If header contains Vary:[Accept-Encoding] param that means that content was compressed
	// resp.Body returns *http.http2gzipReader Reader type
	// Reader get source data and put it in transfer buffer, after returns pointer to this transfer buffer
	// ioutil.ReadAll read data from transfer buffer and returns byte slice
	// body, err := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// if err != nil {
	// 	fmt.Println("Can not get body from response:", err)
	// }
	// fmt.Printf("%s", body)
	parse, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Can not parse response:", err)
	}
	// fmt.Println(parse)
	for _, val := range readlinks(nil, parse) {
		fmt.Println(val)
	}
}
