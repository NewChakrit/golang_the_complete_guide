package main

import (
	"fmt"
	"math"
)

var investmentAmount, years, expectedReturnRate float64

const inflationRate = 2.5

func main() {
	//var (
	//	investmentAmount   float64 = 1000
	//	expectedReturnRate         = 5.5
	//	years              float64 = 10
	//)

	//var (
	//	investmentAmount   float64
	//	years              float64
	//	expectedReturnRate float64
	//)

	//expectedReturnRate := 5.5

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Years: ")
	fmt.Scan(&years)

	//futureValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	//futureRealValue := calculateFutureRealValue(futureValue, inflationRate, years)

	fv, rfv := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	// Outputs Information
	// .2f print 2 digits of float, can chane number between "." and "f" to print how many digits you want.

	formattedFV := fmt.Sprintf("Future Values: %.2f\n", fv)
	// If string tax is too long for line, we can use ``
	formattedRFV := fmt.Sprintf("Future Value (addjusted for inflation): %.3f\n", rfv)

	//fmt.Printf("Future Value: %.2f\n", futureValue)
	fmt.Print(formattedFV)
	fmt.Print(formattedRFV)

	//myCal()
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(invAmount, expReturnRate, years float64) (fv, rfv float64) {
	fv = invAmount * math.Pow(1+expReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)

	return fv, rfv
	//return <= can use only return
}

//func calculateFutureValue(invAmount, expReturnRate, years float64) float64 {
//	return invAmount * math.Pow(1+expReturnRate/100, years)
//}

func calculateFutureRealValue(futureValue, inflationRate, years float64) float64 {
	return futureValue / math.Pow(1+inflationRate/100, years)
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
