package main

import "fmt"

func main() {
	var input string
	var input2 string
	fmt.Print("Enter the First String:")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("please enter valid String")
	}
	fmt.Print("Enter the Second String:")
	_, err2 := fmt.Scan(&input2)
	if err2 != nil {
		fmt.Println("please enter valid String")
	}
	result := input + " " + input2
	fmt.Println("Concatenation ", result)
}
