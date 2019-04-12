package main

import "fmt"

// &arr int - get adress of the variable
// *arr int - value from this address
// arr *int - pointer to a number

func sum(arr *[]int) int {
	var sum int
	for _, f := range *arr {
		sum += f
	}
	return sum
}

func subst(arr *[]int) int {
	sli := *arr
	sub := sli[0]
	for _, f := range sli[1:] {
		sub -= f
	}
	return sub
}
func mult(arr *[]int) int {
	sli := *arr
	mult := sli[0]
	for _, f := range sli[1:] {
		mult *= f
	}
	return mult
}
func div(arr *[]int) int {
	sli := *arr
	div := sli[0]
	for _, f := range sli[1:] {
		div /= f
	}
	return div
}

type countfunctions func(*[]int) int

func count(arr *[]int, a ...countfunctions) []int {
	res := []int{}
	for _, f := range a {
		res = append(res, f(arr))
	}
	return res
}

func main() {
	arr := []int{55, 151, 45, 45, 46} // Slice of ints
	acts := count(&arr, sum, subst, mult, div)
	fmt.Println(acts)
}
