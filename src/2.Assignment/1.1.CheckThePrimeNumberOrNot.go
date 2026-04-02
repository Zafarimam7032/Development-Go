package main

import "fmt"

func main() {
	var input int
	fmt.Print("Enter number: ")
	_, error := fmt.Scan(&input)
	if error != nil {
		fmt.Println("please enter valid Number")
	}
	checkPrimeNumberOrNumber(input)

}

func checkPrimeNumberOrNumber(number int) {
	if number == 2 {
		fmt.Println("given Number is Prime Number")
	} else {
		check := true
		for i := 2; i <= number/2; i++ {
			if number%i == 0 {
				check = false
				break
			}
		}
		if check {
			fmt.Println("given Number is Prime Number")
		} else {
			fmt.Println("given Number is not Prime Number")
		}
	}

}
