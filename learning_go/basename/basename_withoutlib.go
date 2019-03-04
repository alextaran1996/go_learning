package main

import (
	"fmt"
	"os"
)

func removeslash(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func processinput() {
	for _, val := range os.Args[1:] {
		basename := removeslash(val)
		fmt.Println(basename)
	}
}

func main() {
	processinput()
}
