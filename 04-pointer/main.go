package main

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int
	agePointer = &age

	//agePointer := &age

	fmt.Println("Age:", *agePointer) // * ยกเลิก pointer

	getAdultYears(agePointer)
	fmt.Println(age)
}

func getAdultYears(age *int) {
	*age = *age - 18
}
