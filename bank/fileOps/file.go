package file

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteBalanceToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)

	os.WriteFile(fileName, []byte(valueText), 0644)
}

func Choice() int {
	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)
	return choice
}

func GetBalanceFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}

	valueText := string(data)
	balance, _ := strconv.ParseFloat(valueText, 64)
	return balance, nil
}
