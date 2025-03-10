package prices

import (
	"awesomeProject/conversion"
	"awesomeProject/iomanager"
	"fmt"
)

type TaxIncludedPriceJob struct {
	IOManager        iomanager.IOManager `json:"-"`
	TaxRate          float64             `json:"tax_rate"`
	inputPrices      []float64           `json:"input_prices"`
	TaxIncludedPrice map[string]string   `json:"tax_included_price"`
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}
	prices, errr := conversion.StringToFloat(lines)
	if errr != nil {
		fmt.Println("Reading file error")
		fmt.Println(errr)
		return err
	}
	job.inputPrices = prices
	return nil
}
func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}
	result := make(map[string]string)
	for _, price := range job.inputPrices {
		TaxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.f2", TaxIncludedPrice)
	}
	job.TaxIncludedPrice = result
	err = job.IOManager.WriteResult(job)
	if err != nil {
		return err
	}
	return nil
}
func NewTaxIncludedPriceJob(fm iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   fm,
		inputPrices: []float64{},
		TaxRate:     taxRate,
	}
}
