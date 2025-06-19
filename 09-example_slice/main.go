package main

import "fmt"

func main() {
	// 1)
	fmt.Println(" 1)")
	hobbies := [3]string{"reading a book", "listening a song", "drinking a cup of coffee"}
	fmt.Println(hobbies)
	fmt.Println("-----------------------")

	// 2)
	fmt.Println(" 2)")
	hobbies[0] = "running"
	fmt.Println(hobbies[0])
	//fmt.Println(hobbies[1:])
	fmt.Println(hobbies)
	fmt.Println("-----------------------")

	// 3)
	fmt.Println(" 3)")
	newHobbies := hobbies[:1]
	fmt.Println(newHobbies)
	newHobbies = append(newHobbies, hobbies[1])
	fmt.Println(newHobbies)
}
