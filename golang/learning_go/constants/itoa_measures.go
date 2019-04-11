package main

import "fmt"

const (
	_ = uint64(1) << (10 * iota)
	Kb
	Mb
	Gb
	Tb
	Pb
	Eb
)

func main() {
	fmt.Println(Eb)
}
