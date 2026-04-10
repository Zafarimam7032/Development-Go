package main

type AccountErrorHandling struct {
	message string
	date    string
}

type BankAccount struct {
	AccountHolderName string
	AccountId         int64
	AccountType       string
	OpeningBalance    float64
}
