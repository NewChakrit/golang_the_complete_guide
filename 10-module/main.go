package main

import "fmt"

type transfromFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transfromerFn1 := getTransformerFunction(&numbers)
	transfromerFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transfromerFn1)
	moreTransformNumbers := transformNumbers(&moreNumbers, transfromerFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformNumbers)
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

// func getTransformerFunction() func(int) int {
func getTransformerFunction(numbers *[]int) transfromFn {
	if (*numbers)[0] == 1 { // ใช้ ()เพราะต่างโคตรสร้าง
		return double
	} else {
		return triple
	}
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
