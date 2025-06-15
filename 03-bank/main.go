package main

import (
	"fmt"
	"os"
	"strconv"
)

var accountBalanceFile = "balance.txt"

func writeBalacneTofile(balance float64) {
	balanceStr := fmt.Sprint(balance)
	err := os.WriteFile(accountBalanceFile, []byte(balanceStr), 0644)
	if err != nil {
		fmt.Print(err)
	}
}

func readBalanceFromfile() float64 {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		fmt.Print(err)
	}

	balStr := string(data)
	balance, err := strconv.ParseFloat(balStr, 64)
	if err != nil {
		fmt.Print(err)
	}

	return balance
}

func main() {
	balance := readBalanceFromfile()

	fmt.Println("Welcome to Go Bank!")

	var choice, transaction int
	for {
		if transaction > 0 {
			fmt.Println("What do you want to do next?")
		} else {
			fmt.Println("What do you want to do?")
		}
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

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
		fmt.Println("Invalid amount. Must be greater than 0.\n")
		return balance
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
		fmt.Println("Invalid amount. Must be greater than 0.\n")
		return balance
	} else {
		balance -= withdraw
		writeBalacneTofile(balance)
		return balance
	}
}
