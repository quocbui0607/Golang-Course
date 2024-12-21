package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	prices "example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, 4)
	errChans := make([]chan error, 4)

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errChans[index] = make(chan error)
		fm := filemanager.NewFileManager("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.NewCMDManager()
		priceJob := prices.NewTaxIncludedPricesJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errChans[index])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}

	for index := range taxRates {
		select {
		case err := <-errChans[index]:
			if err != nil {
				fmt.Println(err)
			}

		case <-doneChans[index]:
			fmt.Println("Done")
		}
	}
}
