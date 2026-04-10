package main

type BankAccountOperation interface {
	initAccounts() []BankAccount
	withdraw(withdrawAmount float64) AccountErrorHandling
	Deposit(depositAmount float64)
	BalanceCheck()
}
