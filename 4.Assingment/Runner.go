package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var banksAccounts = make([]BankAccount, 0)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to Bank ")
	for {
	O:
		var acountHolderName string
		var accountNumber string
		var accountType string
		var openingBalance float64
		fmt.Println("Please enter Following details... ")
		fmt.Print("Account Holder Name : ")
		scanner.Scan()
		acountHolderName = scanner.Text()
		fmt.Print("Account Number : ")
	Y:
		_, err1 := fmt.Scan(&accountNumber)
		if err1 != nil {
			fmt.Print("please enter valid account number : ")
			goto Y
		}
		fmt.Print("Account Type : ")
	Z:
		_, err2 := fmt.Scan(&accountType)
		if err2 != nil {
			fmt.Print("please enter valid account type : ")
			goto Z
		}
		fmt.Print("Opening Balance : ")
	H:
		scanner.Scan()
		openingBalance, err4 := strconv.ParseFloat(scanner.Text(), 32)
		if err4 != nil {
			fmt.Print("please enter valid opening balance : ")
			goto H
		}
		banksAccount := BankAccount{AccountHolderName: acountHolderName, AccountNumber: accountNumber, AccountType: accountType, OpeningBalance: openingBalance}
		banksAccounts = append(banksAccounts, banksAccount)
		fmt.Println("Do you want enter another Account details ... then Press Y other wise Any Chars")
		var operationType string
		fmt.Scan(&operationType)
		if strings.EqualFold(operationType, "Y") {
			goto O
		}
		break
	}

	for {
		fmt.Println("Welcome to Bank....")
	V:
		var accountNumber string
		fmt.Println("Please enter Account Number : ")

		fmt.Scan(&accountNumber)
		for _, bankAccount := range banksAccounts {
			if strings.EqualFold(bankAccount.AccountNumber, accountNumber) {
				var option string
				fmt.Println("please enter Operation Type Number...")
			X:
				fmt.Println("1.Deposite Balance")
				fmt.Println("2.Withdraw Balance")
				fmt.Println("3.Accounts Information")
				fmt.Scan(&option)
				switch option {
				case "1":
					{
						var depositedBalance float64
						fmt.Print("Please Enter Account Deposited Balance : ")
						fmt.Scan(&depositedBalance)
						bankAccount.Deposit(depositedBalance)
						if oprtationControll() {
							goto X
						}
					}
				case "2":
					{
						var withdrawBalance float64
						fmt.Print("Please Enter Account Withdrew Balance : ")
						fmt.Scan(&withdrawBalance)
						bankAccount.Withdrew(withdrawBalance)
						if oprtationControll() {
							goto X
						}
					}
				case "3":
					{
						bankAccount.AccountDisplay()
						if oprtationControll() {
							goto X
						}
					}
				default:
					{
						fmt.Println("please enter valid Operation Type : ")
						goto X
					}
				}
			} else {
				break
			}
		}
		goto V
	}

	fmt.Println(banksAccounts)
}
func oprtationControll() bool {
	fmt.Println("if you want to Perform other Operation If Yes then Press Y other wise Any Chars...")
	var optionType string
	fmt.Scan(&optionType)
	if strings.EqualFold(optionType, "Y") {
		return true
	}
	return false
}
