package main

import (
	"fmt"

	filemanager "github.com/guetzan/price-calculator/file-manager"
	"github.com/guetzan/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.10, 0.15}
	doneChannels := make([]chan bool, len(taxRates))
	errChannels := make([]chan error, len(taxRates))

	for index, rate := range taxRates {
		fileManager := filemanager.New("prices.txt", fmt.Sprintf("result%.0f.json", rate*100))
		// cmdManager := cmdmanager.New()

		doneChannels[index] = make(chan bool)
		errChannels[index] = make(chan error)

		priceJob := prices.NewTaxIncludedPriceJob(fileManager, rate)
		go priceJob.Process(doneChannels[index], errChannels[index])
	}

	for index := range taxRates {
		select {
		case err := <-errChannels[index]:
			if err != nil {
				fmt.Println(err)
				return
			}

		case <-doneChannels[index]:
			fmt.Println("Job done!")
		}
	}
}
