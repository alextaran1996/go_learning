package main

import (
	"fmt"
	"log"
)

func minmax(values ...int) (max int, min int) { // variative function; it takes unspecified number of int
	if len(values) <= 1 {
		log.Fatal("Minimum number of elements is 2")

	}
	vars := []int{}
	for _, val := range values {
		vars = append(vars, val)
	}
	values = sort(vars)
	min, max = values[0], values[len(values)-1]
	return // emty return.We can use it as we specified var names in returned values
}

func sort(sliceargs []int) []int { // bubble sort
	for i := 0; i < len(sliceargs)-i-1; i++ {
		for j := 0; j < len(sliceargs)-1; j++ {
			if sliceargs[j] > sliceargs[j+1] {
				sliceargs[j], sliceargs[j+1] = sliceargs[j+1], sliceargs[j]
			}
		}
	}
	return sliceargs
}

func main() {
	fmt.Println(minmax(1, 12, 312, 14, 4316, 146))
}
