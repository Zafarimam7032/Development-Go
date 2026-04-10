package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func (b *BankAccount) initAccounts() []BankAccount {
	var accounts []BankAccount
	account1 := BankAccount{AccountHolderName: "zafar Imam", AccountId: 12345, AccountType: "salry", OpeningBalance: 1000.00}
	account2 := BankAccount{AccountHolderName: "James", AccountId: 123, AccountType: "saving", OpeningBalance: 1200.00}
	account3 := BankAccount{AccountHolderName: "jhones", AccountId: 123456, AccountType: "current", OpeningBalance: 1500.00}
	accounts = append(accounts, account1)
	accounts = append(accounts, account2)
	accounts = append(accounts, account3)
	return accounts
}

func (b *BankAccount) withdraw(withdrawAmount float64) AccountErrorHandling {
	if withdrawAmount > b.OpeningBalance {
		return AccountErrorHandling{message: "inSufficient Fund", date: time.DateTime}
	}
	b.OpeningBalance = b.OpeningBalance - withdrawAmount
	return AccountErrorHandling{message: "SuccessFully Withdraw", date: time.DateTime}
}

func (b *BankAccount) Deposit(depositAmount float64) {
	b.OpeningBalance = b.OpeningBalance + depositAmount
	fmt.Println("Successfully Deposit")
}
func (b *BankAccount) BalanceCheck() {
	fmt.Println("Balance Remaining : ", b.OpeningBalance)
}

func main() {

	bankAccount := BankAccount{}

	var bankAccountOperation BankAccountOperation = &bankAccount
	accounts := bankAccountOperation.initAccounts()
	scanner := bufio.NewScanner(os.Stdin)
	for {
	X:
		fmt.Println("Welcome Bank Operation ...")
		var accountNumber int64
		fmt.Println("please enter Account : ")
		check := true
		scanner.Scan()
		accountNumber, err6 := strconv.ParseInt(scanner.Text(), 10, 64)

		if err6 != nil {
			fmt.Print("please enter valid account Number")
		}

		for _, val := range accounts {

			if val.AccountId == accountNumber {
			O:
				check = false
				fmt.Println("please Enter the Option ...")
				fmt.Println("1.withdraw")
				fmt.Println("2.Deposit")
				fmt.Println("3.Check Balance")
				var option string
				_, err := fmt.Scan(&option)
				if err != nil {
					fmt.Println("please enter valid option")
				}
				if strings.EqualFold(option, "1") {
					fmt.Print("please enter amount to withdraw : ")
					var amount float64
				A:
					scanner.Scan()
					amount, err2 := strconv.ParseFloat(scanner.Text(), 64)
					if err2 != nil {
						fmt.Print("please enter valid amount : ")
						goto A
					}
					info := val.withdraw(amount)
					fmt.Println(info)

				} else if strings.EqualFold(option, "2") {
					fmt.Print("please enter for deposit : ")
					var amount float64
				B:
					scanner.Scan()
					amount, err3 := strconv.ParseFloat(scanner.Text(), 64)
					if err3 != nil {
						fmt.Print("please enter valid amount : ")
						goto B
					}
					val.Deposit(amount)

				} else if strings.EqualFold(option, "3") {
					val.BalanceCheck()

				}
				fmt.Println("If you want to Perform other Operation please Type Y otherwise Any character ")
				var operation string
				fmt.Scan(&operation)
				if strings.EqualFold(operation, "y") {
					goto O
				}
				break
			}

		}
		if check {
			fmt.Print("please enter valid account Number : ")
		}

		goto X
	}

}
