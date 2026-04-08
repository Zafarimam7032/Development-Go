package main

import "fmt"

func (bankaccount *BankAccount) Withdrew(amount float64) {
	if amount > bankaccount.OpeningBalance {
		fmt.Println("insufficient balance")
	}
	withrownAmmount := bankaccount.OpeningBalance - amount
	fmt.Println("Remaining balance ", withrownAmmount)
	bankaccount.OpeningBalance = withrownAmmount
}

func (bankaccount *BankAccount) Deposit(amount float64) {
	depositAmount := bankaccount.OpeningBalance + amount
	fmt.Println("Total amount ", depositAmount)
	bankaccount.OpeningBalance = depositAmount
}

func (account BankAccount) AccountDisplay() {
	fmt.Println("\nAccount Information....")
	fmt.Println("1.AccountHolderName: ", account.AccountHolderName)
	fmt.Println("2.AccountNumber: ", account.AccountNumber)
	fmt.Println("3.AccountType: ", account.AccountType)
	fmt.Println("4.AccountBalance: ", account.OpeningBalance)

}
