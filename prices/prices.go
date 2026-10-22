package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/guetzan/price-calculator/converter"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	Prices            []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	for _, price := range job.Prices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		job.TaxIncludedPrices[fmt.Sprintf("%.2f", price)], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", taxIncludedPrice), 64) //coisa horrorosa (temporária)
	}

	fmt.Println(job.TaxIncludedPrices)
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("Failed to open file.")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)

	var content []string

	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Error at reading the file.")
		fmt.Println(err.Error())
		file.Close()
		return
	}

	file.Close()

	prices, err := converter.StringsToFloat(content)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	job.Prices = prices
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:           taxRate,
		Prices:            []float64{54.99, 99.90, 8.50},
		TaxIncludedPrices: make(map[string]float64),
	}
}
