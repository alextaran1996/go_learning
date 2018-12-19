package main

import "fmt"

func hello_world(name string) {
	fmt.Printf("Hello,%s\n", name)
}

func main() {
	hello_world("Ken Thompson")
}
