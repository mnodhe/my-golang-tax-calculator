package prices

import (
	"awesomeProject/conversion"
	"awesomeProject/filemanager"
	"fmt"
)

type TaxIncludedPriceJob struct {
	IOManager        filemanager.FileManager `json:"-"`
	TaxRate          float64                 `json:"tax_rate"`
	inputPrices      []float64               `json:"input_prices"`
	TaxIncludedPrice map[string]string       `json:"tax_included_price"`
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	prices, errr := conversion.StringToFloat(lines)
	if errr != nil {
		fmt.Println("Reading file error")
		fmt.Println(errr)
		return
	}
	job.inputPrices = prices
	return
}
func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.inputPrices {
		TaxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.f2", TaxIncludedPrice)
	}
	job.TaxIncludedPrice = result
	job.IOManager.WriteResult(job)
}
func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   fm,
		inputPrices: []float64{},
		TaxRate:     taxRate,
	}
}
