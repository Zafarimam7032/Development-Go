package main

import "fmt"

func main() {
	var input int
	fmt.Print("Enter number: ")
	val, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("please enter valid Number")
	} else {
		fmt.Println(val)
	}

	CheckPalindromeNumber(input)

}
func CheckPalindromeNumber(number int) {
	var result int
	var num = number
	for number > 0 { //231==>1,,23
		rem := number % 10
		number /= 10
		result = result*10 + rem
	}
	if result == num {
		fmt.Println("given Number is  Palindrome")
	} else {
		fmt.Println("given Number is not Palindrome")
	}
}
