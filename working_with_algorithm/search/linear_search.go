package main

import (
	"fmt"
	"time"
)

func linearsearch(slice []int, search int) ([]int, time.Duration) { // Linear search
	start := time.Now()    // Get current time
	searchindex := []int{} // Create slice for index values
	for ind, val := range slice {
		if val == search {
			searchindex = append(searchindex, ind) // If element was found add it index in searchindex slice
		}
	}
	elapsed := time.Since(start) // Get time elapsed from the beginnig of the search
	return searchindex, elapsed  // Return slice with indexes and time

}

func main() {
	unsorted := []int{12, 32, 5443, 32, -21, 312, -1562, 1454, 135, -3512, -1562, 35235, 235, 3, 21, 415, 325, 642}
	var search = -1562
	res, exectime := linearsearch(unsorted, search)
	if len(res) == 0 {
		fmt.Println(search, "doesn't exist in the specified slice.\nExecution time:", exectime)
	} else {
		fmt.Println(search, "element is/are on the", res, "position/s.\nExecution time:", exectime)
	}
}
