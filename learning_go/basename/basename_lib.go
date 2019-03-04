package main

import (
	"fmt"
	"os"
	"strings"
)

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func proccessinput() {
	for _, val := range os.Args[1:] {
		base := basename(val)
		fmt.Println(base)
	}

}

func main() {
	proccessinput()
}
