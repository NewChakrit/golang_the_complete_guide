package main

import (
	"fmt"

	"bank/fileops"
)

var accountBalanceFile = "balance.txt"

func main() {
	balance, err := fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		//fmt.Printf("[Error] %v", err)
		//fmt.Print("\nSet default balance to 1,000\n")
		//fmt.Print("------------------------------\n")

		fmt.Print(err)
	}

	fmt.Println("Welcome to Go Bank!")

	var choice, transaction int
	for {
		presentOptions(transaction)

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		//wantsCheckBalance := choice == 1 // <= can you this return bool
		//if wantsCheckBalance {
		//
		//}

		/* If we use break in switch case. That mean break in switch case not break for loop? */

		switch choice {
		case 1:
			fmt.Println("------------------------------")
			fmt.Printf("Your balance: %v\n", balance)
			fmt.Println("------------------------------")
			transaction++
			continue
		case 2:
			balance = calculateDeposit(balance)

			fmt.Println("------------------------------")
			fmt.Printf("Balance updated! New amount: %v\n", balance)
			fmt.Println("------------------------------")
			transaction++
			continue
		case 3:
			balance = calculateWithDraw(balance)

			fmt.Println("------------------------------")
			fmt.Printf("Balance updated! New amount: %v\n", balance)
			fmt.Println("------------------------------")
			transaction++
			continue
		case 4:
			fmt.Println("-------------------------------")
			fmt.Println("Thank you for using our service :)")
			fmt.Println("-------------------------------")
			return
		default:
			fmt.Println("-------------------------------")
			fmt.Println("Please select the required option.")
			fmt.Println("-------------------------------")
			transaction++
		}
	}

}

func calculateDeposit(balance float64) float64 {
	var deposit float64
	fmt.Print("How much you want to deposit?: ")
	fmt.Scan(&deposit)

	if deposit <= 0 {
		//fmt.Println("Invalid amount. Must be greater than 0.\n")
		//return balance

		panic("[ERROR] deposit must greatest than 0.")
	}

	balance += deposit

	fileops.WriteFloatTofile(accountBalanceFile, balance)
	return balance
}

func calculateWithDraw(balance float64) float64 {
	var withdraw float64
	fmt.Print("How much you want to withdrawal?: ")
	fmt.Scan(&withdraw)
	if balance < withdraw {
		fmt.Print("Invalid amount. You can not withdraw more than you have.\n")
		fileops.WriteFloatTofile(accountBalanceFile, balance)
		return balance
	} else if withdraw <= 0 {
		//fmt.Println("Invalid amount. Must be greater than 0.\n")
		//return balance

		panic("[ERROR] withdraw must greatest than 0.")
	} else {
		balance -= withdraw
		fileops.WriteFloatTofile(accountBalanceFile, balance)
		return balance
	}
}
