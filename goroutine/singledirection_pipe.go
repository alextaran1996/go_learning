package main

import "fmt"

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func square(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printler(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	in := make(chan int)
	out := make(chan int)
	go counter(out)
	go square(in, out)
	printler(in)

}
