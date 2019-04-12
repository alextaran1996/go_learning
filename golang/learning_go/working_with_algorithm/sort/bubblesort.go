package main

import (
	"fmt"
	"time"
)

// Algorithm
// For every element in the array compare neighboring elements.Swap them if left element bigger then right.
// Repeat this action len(array) times

// Average  Sorting time is 730ns
func bubblesort(arr []int) ([]int, string) {
	start := time.Now()
	for i := 0; i < len(arr)-i-1; i++ { // Using len(arr)-i-1, because every cycle the biggest number appeared at the end of the array, therefor we don't need compare it in the next cycles
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	elapsed := "\nSorting time is :" + time.Since(start).String()
	return arr, elapsed
}

func main() {
	unsorted := []int{12, 32, 5443, 32, -21, 312, -1562, 1454, 135, -3512, -1562, 35235, 235, 3, 21, 415, 325, 642}
	fmt.Println(bubblesort(unsorted))
}
