package main

import "fmt"

func main() {
	var countryCode string
	fmt.Print("please enter Country code : ")
	_, err := fmt.Scan(&countryCode)
	if err != nil {
		return
	}
	var amount int
	fmt.Print("please enter amount: ")
	_, error1 := fmt.Scan(&amount)
	if error1 != nil {
		return
	}
	currencyDisplay(amount, countryCode)

}

func currencyDisplay(amount int, countryCode string) {

	currencySymbols := map[string]string{
		"USD": "$",
		"EUR": "€",
		"GBP": "£",
		"INR": "₹",
		"JPY": "¥",
	}

	sym, ok := currencySymbols[countryCode]
	if ok {
		fmt.Printf("%s %d\n", sym, amount)
	}

}
