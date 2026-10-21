package main

import (
	"github.com/guetzan/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.10, 0.15}

	for _, rate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(rate)
		priceJob.Process()
	}
}
