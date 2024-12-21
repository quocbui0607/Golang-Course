package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	outputText("Enter investmentAmount: ")
	fmt.Scan(&investmentAmount)

	outputText("Enter expectedReturnRate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Enter years: ")
	fmt.Scan(&years)

	fmt.Println("-------- Result --------")

	futureValue, futureRealValue := calculateFeatureValues(investmentAmount, expectedReturnRate, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFeatureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	frv := fv / math.Pow((1+inflationRate/100), years)
	return fv, frv
}
