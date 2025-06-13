package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount = 1000.
	var expectedReturnRate = 5.5
	var years = 10

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	fmt.Println(futureValue) // 1708.1444583535929

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
