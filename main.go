package main

import (
	"awesomeProject/cmdmanager"
	"fmt"
	//"awesomeProject/filemanager"
	"awesomeProject/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	for _, taxRate := range taxRates {
		//fm := filemanager.FileManager{InputFilePath: "prices.txt", OutputFilePath: fmt.Sprintf("result_%.0f.json", taxRate*100)}
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("Could not process job")
		}
	}
}

// watched 139
