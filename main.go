package main

import "fmt"

func main() {
	prices := []float64{13, 21, 34}
	taxRates := []float64{0, 0.07, 0.10, 0.15}
	taxIncludedPrices := make(map[float64][]float64)

	for _, rate := range taxRates {
		newPrices := make([]float64, len(prices))

		for index, price := range prices {
			newPrices[index] = price * (1 + rate)
		}

		taxIncludedPrices[rate] = newPrices
	}

	for key, value := range taxIncludedPrices {
		fmt.Printf("%v%%: %v\n", key, value)
	}
}
