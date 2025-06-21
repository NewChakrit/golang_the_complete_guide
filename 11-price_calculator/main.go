package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	//taxRates := []float64{0, 0.07, 0.1, 0.15}
	//
	//for _, taxRate := range taxRates {
	//	priceJob := prices.NewTaxIncludedPriceJob(taxRate)
	//	priceJob.Process()
	//}

	price, err := os.ReadFile("prices.txt")
	if err != nil {
		fmt.Println(errors.New("Can not read prices from prices.txt"))
	}

	fmt.Println(price)
}
