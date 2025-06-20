package main

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func main() {
	// 1)
	fmt.Println(" 1)")
	hobbies := [3]string{"reading a book", "listening a song", "drinking a cup of coffee"}
	fmt.Printf("1.1) %v\n", hobbies)
	fmt.Println("-----------------------")

	// 2)
	fmt.Println(" 2)")
	fmt.Printf("2.1) %v\n", hobbies[0])
	secondList := hobbies[1:]
	fmt.Printf("2.2) %v\n", secondList)
	fmt.Println("-----------------------")

	// 3)
	fmt.Println(" 3)")
	mainHobbies := hobbies[:2]
	fmt.Printf("3.1) %v\n", mainHobbies)

	//newHobbiesTwo := []string{hobbies[0], hobbies[1]}
	//newHobbiesTwo = append(newHobbiesTwo, hobbies[0], hobbies[1])
	//fmt.Printf("3.2) %v\n", newHobbiesTwo)
	fmt.Println("-----------------------")

	// 4)
	mainHobbies = mainHobbies[1:3]
	fmt.Printf("4) %v\n", mainHobbies)
	fmt.Println("-----------------------")

	// 5)
	courseGoals := []string{"Learn Go", "Learn Java"}
	fmt.Printf("5) %v\n", courseGoals)
	fmt.Println("-----------------------")

	// 6)
	courseGoals[1] = "Learn JavaScript"
	courseGoals = append(courseGoals, "Learn HTML")
	fmt.Printf("6) %v\n", courseGoals)
	fmt.Println("-----------------------")

	// 7)
	products := []Product{
		{
			title: "Learn Go",
			id:    1,
			price: 399.99,
		},
		{
			title: "Learn Javascript",
			id:    2,
			price: 249.5,
		},
	}

	products = append(products, Product{"Learn CSS", 3, 199.25})
	fmt.Printf("7) %v\n", products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
