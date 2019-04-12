package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// SiteURL is Url of the site that provides json as a response
const SiteURL = "http://reqres.in/api/users/3"

type Poster struct {
	Data *Params // Pointer to another structure that includes data fields
}

type Params struct {
	ID        int    // Remeber! The name of the field descriptor should be surronded by "".Example `json:"name_of_your_field"`
	Firstname string `json:"first_name"` // Field descriptor.Required to correlate fields from JSON file with your structure
	Lastname  string `json:"last_name"`  // If file descriptor don't correspond to name in JSON file, this field will be ignored
}

func main() {
	resp, err := http.Get(SiteURL) // Make a Get request
	if err != nil {
		fmt.Println("Resource is unavaliable")
	}
	defer resp.Body.Close()               // Release resources related to the request
	if resp.StatusCode != http.StatusOK { // Check response code is 200
		fmt.Println("Resourse is unavaliable")
	}
	// if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
	// 	resp.Body.Close()

	// }	// Alternative method to decode response
	jsontext, _ := ioutil.ReadAll(resp.Body)
	var res Poster
	errors := json.Unmarshal(resp, &res) // Parse data from response
	if errors != nil {
		log.Fatal(errors)
	}
	fmt.Println(res.Data)

}
