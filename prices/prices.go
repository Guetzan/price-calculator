package prices

import (
	"fmt"

	"github.com/guetzan/price-calculator/converter"
)

type TaxIncludedPriceJob struct {
	IOManager         IOManager         `json:"-"`
	TaxRate           float64           `json:"tax_rate"`
	Prices            []float64         `json:"prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}

type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(data any) error
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errChan chan error) {
	err := job.LoadData()

	if err != nil {
		errChan <- err
		return
	}

	result := make(map[string]string)

	for _, price := range job.Prices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	err = job.IOManager.WriteResult(job)

	doneChan <- true
}

func (job *TaxIncludedPriceJob) LoadData() error {
	content, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := converter.StringsToFloat(content)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	job.Prices = prices
	return nil
}

func NewTaxIncludedPriceJob(ioManager IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:         ioManager,
		TaxRate:           taxRate,
		Prices:            []float64{54.99, 99.90, 8.50},
		TaxIncludedPrices: make(map[string]string),
	}
}
