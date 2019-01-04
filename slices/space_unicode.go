package main

import (
	"fmt"
	"unicode"
)

func rotate(strs []byte) []byte {
	st := strs[:0]
	for i, s := range strs {
		if unicode.IsSpace(rune(s)) {
			if unicode.IsSpace(rune(strs[i-1])) && i > 0 {
				continue
			} else {
				st = append(st, ' ')
			}

		} else {
			st = append(st, s)
		}

	}
	return st
}

func main() {
	s := []byte("This is         simple string")
	fmt.Printf("%q\n", string(rotate(s)))
	fmt.Printf("%q\n", s)
}
