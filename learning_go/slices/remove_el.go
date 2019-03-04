package main

import "fmt"

func removeel(sl []int, i int) []int {
	copy(sl[i:], sl[i+1:])
	return sl[:len(sl)-1]
}

func main() {
	a := []int{2, 3, 54, 4, 2, 6}
	fmt.Println(removeel(a, 1))
	// b := []int{2, 3, 4, 5}
	// c := []int{1, 1, 1}
	// copy(b[2:], c[0:1])
	// fmt.Println(b)
}
