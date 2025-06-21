package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		//fmt.Println("Tax Rate:", val)

		taxIncludedPrices := make([]float64, len(prices))
		for i, p := range prices {

			taxIncludedPrices[i] = p * (1 + taxRate)
		}

		result[taxRate] = taxIncludedPrices
	}

	for k, v := range result {
		fmt.Printf("Tax Rate: %.2f\n", k)
		fmt.Printf("Price: %.2f\n", v)
	}
}
