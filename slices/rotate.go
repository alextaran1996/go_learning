package main

import "fmt"

func rotate(strs []string) []string {
	w := 0
	for _, s := range strs {
		if strs[w] == s {
			continue
		}
		w++
		strs[w] = s
	}
	return strs[:w+1]
}

func main() {
	s := []string{"a", "a", "b", "c", "c", "c", "d", "d", "e"}
	fmt.Println(rotate(s))
}
