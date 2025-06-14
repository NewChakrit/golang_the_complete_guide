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
	var investmentAmount float64
	years, expectedReturnRate := 10.0, 5.5

	fmt.Scan(&investmentAmount)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Printf("Future Value: %v\n", futureValue) // 1708.1444583535929
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
