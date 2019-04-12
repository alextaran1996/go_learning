package main

import (
	"fmt"
	"strings"
)

func expand(s string, change string, f func(string) string) string { // Using function as a parameter(first-class function)
	return strings.Replace(s, "$foo", f(change), -1)
}

func main() {
	f := func(str string) string { // f has function type
		return strings.ToUpper(str)
	}
	inputstr := "Yo, Big Shaq, the one and only\nMan's not hot, never hot\n$foo\nBoom"
	change := "Skrrat, skidi-kat-kat"
	as := expand(inputstr, change, f)
	fmt.Println(as)

}

// func add(a int, b int) int {
// 	return a + b
// }

// func susbt(a int, b int) int {
// 	return a - b
// }

// func chooseact(a int, b int, f func(int, int) int) int {
// 	res := f(a, b)
// 	return res
// }

// func main() {
// 	fmt.Println(add(1, 3))
// 	fmt.Println(susbt(1, 3))
// 	fmt.Println(chooseact(1, 56, add))
