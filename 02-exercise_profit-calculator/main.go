package main

import (
	"errors"
	"fmt"
	"os"
)

//var revenue, expenses, taxRate float64

/* Goals
	1) Validate user input
		=> Show error message & exit if invalid input is provided
		- No negative numbers
		- Not 0
    2( Store calculated results into file
*/

func writeFinancialToFile(ebt, profit, ratio float64) {
	financialStr := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
	err := os.WriteFile("Financial.txt", []byte(financialStr), 0644)
	if err != nil {
		panic("Cannot write financial results into file.")
	}
}

func main() {

	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		panic(err)
	}

	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		panic(err)
	}

	taxRate, err := getUserInput("Tex Rate: ")
	if err != nil {
		panic(err)
	}

	fmt.Println("---------------------")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	fmt.Println("EBT = ", ebt)
	fmt.Println("Profit = ", profit)
	fmt.Printf("Ratio = %.2f\n", ratio)

	writeFinancialToFile(ebt, profit, ratio)
}

func getUserInput(text string) (float64, error) {
	var value float64

	fmt.Print(text)
	fmt.Scan(&value)
	if value <= 0 {
		return 0, errors.New("[Error] Value must greatest than 0.")
	}

	return value, nil
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
