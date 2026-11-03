package main

import (
	"fmt"

	filemanager "github.com/guetzan/price-calculator/file-manager"
	"github.com/guetzan/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.10, 0.15}

	for _, rate := range taxRates {
		fileManager := filemanager.New("prices.txt", fmt.Sprintf("result%.0f.json", rate*100))
		// cmdManager := cmdmanager.New()

		priceJob := prices.NewTaxIncludedPriceJob(fileManager, rate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Failed to process job.")
			fmt.Println(err.Error())
			return
		}
	}
}
