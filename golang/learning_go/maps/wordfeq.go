package main

import (
	"bufio"
	"fmt"
	"os"
)

func proc_file(file string) map[string]int {
	findmap := make(map[string]int)
	f, err := os.Open(file) // Open file for reading
	if err != nil {
		fmt.Printf("%s : no such file or directory\n", file)
		os.Exit(1)
	} else {
		scanner := bufio.NewScanner(f) // Create scanner for reading text
		scanner.Split(bufio.ScanWords) // Choose function for proccessing input.Split incoming text by words
		for scanner.Scan() {           // Move to the next word until EOF
			word := scanner.Text() // Read a word provided by the current token
			findmap[word]++
		}

		return findmap

	}
	f.Close() // Close file reading
	return findmap

}

func proc_inp() {
	findmap := make(map[string]int)
	for _, val := range os.Args[1:] { // in case of multiple files
		findmap = proc_file(val)          // Call function for counting words
		for key, value := range findmap { // Make readable output
			fmt.Println(key, "\t", value)
		}
	}
}

func main() {
	proc_inp()
}
