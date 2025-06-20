package main

import "fmt"

func main() {
	//list.List()
	//maps.Maps()

	userNames := make([]string, 2, 5)
	// TODO ^ Set default array (if second parameter has value) if not has value, it is slice!.
	// TODO ^ third paramet is capacity
	fmt.Println(len(userNames)) // 2 <= because it default value
	fmt.Println(cap(userNames)) // 5

	userNames[0] = "Meen"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel", "Man", "Mop", "Mo")

	fmt.Println(len(userNames)) // 7 add on default 2
	fmt.Println(cap(userNames)) // 10 // TODO add topup default capacity, EX if default cap 5 and then add value more than 5, cap is 10.
	fmt.Println(userNames)      // [ Meen * Max Manuel]

	// if use
	//name := []string{}
	//name[0] = "John" // TODO <= ERROR!
}
