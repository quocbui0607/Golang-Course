package main

import "fmt"

func main() {
	productNames := [4]string{"A book"}
	prices := []float64{10.99, 9.99, 45.99, 20.0}

	productNames[2] = "A Carpet"

	fmt.Println(prices[3])
	fmt.Println(productNames)

	newPrices := prices[1:2]
	fmt.Println("NewPrice: ", newPrices[:3])

	prices = append(newPrices, 9.99, 12.99, 25.99, 100.10)

	fmt.Println("discountPrices: ", prices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(discountPrices, discountPrices...)

	fmt.Println("Price with discountPrices: ", prices)
	prices[5] = 0
	fmt.Println("discountPrices: ", discountPrices)
	fmt.Println("Price with discountPrices: ", prices)

}
