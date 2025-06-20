package main

import "fmt"

func main() {
	number := []int{1, 2, 3}

	transfromed := transformNumbers(&number, func(num int) int {
		return num * 2
	})
	fmt.Println(transfromed)

	double := createTransformer(2)
	doubled := transformNumbers(&number, double)
	fmt.Println(doubled)

	triple := createTransformer(3)
	tripled := transformNumbers(&number, triple)
	fmt.Println(tripled)

}

func transformNumbers(number *[]int, transform func(int) int) []int {
	//var double []int
	dNumbers := []int{}
	//double := make([]int, len(*number))

	for _, v := range *number {
		dNumbers = append(dNumbers, transform(v))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {

	return func(num int) int {
		return num * factor
	}
}
