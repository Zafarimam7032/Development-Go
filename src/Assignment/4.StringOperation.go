package main

import "fmt"

func main() {
	fmt.Print("Enter a string: ")
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Length of String:", len(input))
	fmt.Println("Last Character of the String: ", input[len(input)-1])
	fmt.Println("Printing Character ASCII...")
	for i := 0; i < len(input); i++ {
		fmt.Println("character : ", string(input[i]), " ASCII : ", input[i])
	}
	fmt.Println("finished ")
	fmt.Print("Revered String: ")
	for i := len(input) - 1; i >= 0; i-- {
		fmt.Print(string(input[i]))
	}

}
