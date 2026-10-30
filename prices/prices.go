package prices

import (
	"fmt"

	"github.com/guetzan/price-calculator/converter"
	filemanager "github.com/guetzan/price-calculator/file-manager"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager
	TaxRate           float64
	Prices            []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.Prices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	job.IOManager.WriteJson(job)
}

func (job *TaxIncludedPriceJob) LoadData() {
	content, err := job.IOManager.ReadFile()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	prices, err := converter.StringsToFloat(content)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	job.Prices = prices
}

func NewTaxIncludedPriceJob(fileManager filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:         fileManager,
		TaxRate:           taxRate,
		Prices:            []float64{54.99, 99.90, 8.50},
		TaxIncludedPrices: make(map[string]string),
	}
}
