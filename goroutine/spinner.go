package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 46
	fibN := fib(n)
	fmt.Printf("\r%d", fibN)
}

func spinner(duration time.Duration) {
	for {
		for _, val := range `-\|/` {
			fmt.Printf("\r%c", val)
			time.Sleep(duration)
		}
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
