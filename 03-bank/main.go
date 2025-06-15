package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var accountBalanceFile = "balance.txt"

func writeBalacneTofile(balance float64) {
	balanceStr := fmt.Sprint(balance)
	err := os.WriteFile(accountBalanceFile, []byte(balanceStr), 0644)
	if err != nil {
		panic(err)
	}
}

func readBalanceFromfile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("[ERROR] Failed to find balance file.\nSet default to balance 1000.\n")
	}

	balStr := string(data)
	balance, err := strconv.ParseFloat(balStr, 64)
	if err != nil {
		return 1000, errors.New("Cannot parse stored balance value.\n")
	}

	return balance, nil
}

func main() {
	balance, err := readBalanceFromfile()
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

	writeBalacneTofile(balance)
	return balance
}

func calculateWithDraw(balance float64) float64 {
	var withdraw float64
	fmt.Print("How much you want to withdrawal?: ")
	fmt.Scan(&withdraw)
	if balance < withdraw {
		fmt.Print("Invalid amount. You can not withdraw more than you have.\n")
		writeBalacneTofile(balance)
		return balance
	} else if withdraw <= 0 {
		//fmt.Println("Invalid amount. Must be greater than 0.\n")
		//return balance

		panic("[ERROR] withdraw must greatest than 0.")
	} else {
		balance -= withdraw
		writeBalacneTofile(balance)
		return balance
	}
}
