package main

import "fmt"

type transfromFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)
}

// func transformNumbers(number *[]int, transform func(int) int) []int {
func transformNumbers(number *[]int, transform transfromFn) []int {
	//var double []int
	dNumbers := []int{}
	//double := make([]int, len(*number))

	for _, v := range *number {
		dNumbers = append(dNumbers, transform(v))
	}

	return dNumbers
}

func doubleNumbers(number *[]int) []int {
	//var double []int
	dNumbers := []int{}
	//double := make([]int, len(*number))

	for _, v := range *number {
		dNumbers = append(dNumbers, double(v))
	}

	return dNumbers
}

func double(num int) int {
	return num * 2
}

func triple(num int) int {
	return num * 3
}
