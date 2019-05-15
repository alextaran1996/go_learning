package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Article is a simple strcture
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles is a slice of Article
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) { // Create example instance of Article
	articles := Articles{
		Article{Title: "Test title", Desc: "Test Description", Content: "Hello world"},
	}
	fmt.Println("Endpoint hit: All Articles Endpoint") // Make output in console
	json.NewEncoder(w).Encode(articles)                // Encode() creates json ,getting data from article -> NewEncoder writes this data in io.Writer
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint Hit") // Fprintf sends its ouput in io.Writers.In this example you will see this word as a response on your request
}

func handleRequests() {
	http.HandleFunc("/", homePage)               // Execute homePage function when server will get request to 'http://yourserver:8081/'
	http.HandleFunc("/articles", allArticles)    // Execute allArticle function when server will get request to 'http://yourserver:8081/articles'
	log.Fatal(http.ListenAndServe(":8081", nil)) // Start web server on localhost:8081 with no default handler
	// log.Fatal is equivalent to Print,but it works after os.Exit(1) function e.g server error
}

func main() {
	handleRequests() // Run web server
}
