package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) { // ResponseWriter is interface used to form a reponse
	fmt.Fprintf(w, "Response") // Form string and send it to writer -> ResponseWriter
}

func main() {
	port := ":8989"                                   // Port for incoming requests
	http.HandleFunc("/", handler)                     // How handle incoming requests
	log.Printf("Started server on localhost%s", port) // Output on which port server started
	log.Fatal(http.ListenAndServe(port, nil))         // If ListenAndServe failed to open specified port log.Fatal will report abouth this problem
}
