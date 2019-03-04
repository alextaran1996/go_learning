package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // Map for storing strings and number of mentions of this strings
	files := os.Args[1:]           // Get arguments from CLI
	if len(files) == 0 {           // If no arguments were provided, read strings from CLI
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files { // Split provided files from CLI
			f, err := os.Open(arg) // Open file for reading and return pointer to the file
			if err != nil {        // Return error if can't find specified file
				fmt.Fprintf(os.Stderr, "diff_file: %v\n", err)
				continue // Processing next file
			}
			countLines(f, counts) // Read file line by line and enter them in map
			f.Close()             // Close file
		}
	}
	for line, n := range counts { // Output number of matches
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}

}
