package main

import (
	"bufio"
	"fmt"
	"os"
)

// You can stop input by pressing Ctr + d
func main() {
	counts := make(map[string]int)      // Create map key-string/value-integer
	input := bufio.NewScanner(os.Stdin) // Read from input stream
	for input.Scan() {                  // Read from stdin and split input by lines
		counts[input.Text()]++ // Check was the line entered previously, if so increase value of this tring(means how many times we get this string from input)
	}
	for line, n := range counts { // Output how may times entered strings were met
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
