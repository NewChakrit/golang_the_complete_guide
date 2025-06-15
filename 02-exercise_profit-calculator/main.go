package main

import "fmt"

//var revenue, expenses, taxRate float64

func main() {

	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tex Rate: ")

	fmt.Println("---------------------")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	fmt.Println("EBT = ", ebt)
	fmt.Println("Profit = ", profit)
	fmt.Printf("Ratio = %.2f\n", ratio)
}

func getUserInput(text string) float64 {
	var value float64
	fmt.Print(text)
	fmt.Scan(&value)

	return value
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}

/*
EX
Ask for revenue, expenses & tax rate (profit)
Calculate earnings before tax (EBT) and earnings after tax(profit)
Calculate ratio (EBT / Profit)
Output EBT, profit and the ratio
*/
