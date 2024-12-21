package main

import (
	"fmt"

	randomData "github.com/Pallinder/go-randomdata"
	file "test.com/bank/fileOps"
)

func main() {
	var accountBalance, err = file.GetBalanceFromFile(balanceFileName)

	if err != nil {
		panic(err)
	}

	for {
		fmt.Println("Reach us ", randomData.PhoneNumber())

		printChoices()

		choice := file.Choice()
		fmt.Println("Your choice:", choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Your deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			file.WriteBalanceToFile(accountBalance, balanceFileName)
		case 3:
			var withdrawalAmount float64

			fmt.Print("Withdrawal amount: ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}
			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			file.WriteBalanceToFile(accountBalance, balanceFileName)

		default:
			fmt.Println("Exit")
			return
		}

	}

}

const balanceFileName = "balance.txt"
