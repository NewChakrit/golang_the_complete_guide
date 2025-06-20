package recursion

import "fmt"

func main() {
	fmt.Println(factorial(5))

}

func factorial(num int) int {
	//result := 1
	//for i := 1; i <= num; i++ {
	//	result *= i
	//}

	if num == 0 {
		return 1
	} // break loop

	return num * factorial(num-1) // loop
}

//factorial of 5 => 5 * 4 * 3 * 2 * 1 = 120
