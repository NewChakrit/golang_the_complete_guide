package main

import "fmt"

func main() {
	var revenue, expenses int64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	fmt.Println("---------------------")

	ebt := float64(revenue - expenses)
	fmt.Println("EBT = ", ebt)

	profit := ebt - (ebt * taxRate / 100)
	fmt.Println("Profit = ", profit)

	ratio := ebt / profit
	fmt.Println("Ratio = ", ratio)
}

/*
EX
Ask for revenue, expenses & tax rate (profit)
Calculate earnings before tax (EBT) and earnings after tax(profit)
Calculate ratio (EBT / Profit)
Output EBT, profit and the ratio
*/
