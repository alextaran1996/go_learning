package main

import (
	"fmt"
	"sync"
)

func main() {
	slice := []string{"a", "b", "c", "d", "e"}
	sliceLength := len(slice)

	var wg sync.WaitGroup
	wg.Add(sliceLength)

	fmt.Println("Running for loop…")

	for i := 0; i < sliceLength; i++ {
		go func(i int) {
			defer wg.Done()
			val := slice[i]
			fmt.Printf("i: %v, val: %v\n", i, val)
		}(i)
	}

	fmt.Println("Doing other stuff")

	wg.Wait()

	fmt.Println("Finished for loop")
}
