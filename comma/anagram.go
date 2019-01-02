package main

import (
	"fmt"
	"os"
	"strings"
)

func anagram(word1 string, word2 string) bool {
	if len(word1) == len(word2) {
		for _, val := range word1 {
			if strings.Count(word1, string(val)) == strings.Count(word2, string(val)) {
				continue
			} else {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func proc_args() {
	word1, word2 := os.Args[1], os.Args[2]
	fmt.Println(anagram(word1, word2))

}

func main() {
	proc_args()
}
