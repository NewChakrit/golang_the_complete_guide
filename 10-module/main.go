package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup(1, 10, 15, 40, -5)    // todo 1 จะเป็น parameter แรก และที่เหลือเป็น parameter ที่ 2
	anotherSum := sumup(1, numbers...) // todo ... = split slice

	fmt.Println(sum)        // 61
	fmt.Println(anotherSum) // 27
}

// func sumup(numbers ...int) int {
func sumup(startingValue int, numbers ...int) int {
	sum := startingValue

	for _, v := range numbers {
		sum += v
	}

	return sum
}
