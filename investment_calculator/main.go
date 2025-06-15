package main

import (
	"fmt"
	"math"
)

func main() {
	//var (
	//	investmentAmount   float64 = 1000
	//	expectedReturnRate         = 5.5
	//	years              float64 = 10
	//)
	const inflationRate = 2.5
	//var (
	//	investmentAmount   float64
	//	years              float64
	//	expectedReturnRate float64
	//)

	var investmentAmount, years, expectedReturnRate float64

	//expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Printf("Future Value: %v\n", futureValue)
	fmt.Printf("Real Future Value: %v\n", futureRealValue)

	//myCal()
}

func myCal() {
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10

	var total = float64(investmentAmount)

	for i := 1; i <= years; i++ {
		total = total + (total*expectedReturnRate)/100
		fmt.Printf("Year %v: %.2f\n", i, total)
	}
}
