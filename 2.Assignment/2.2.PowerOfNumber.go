package main

import (
	"fmt"
	"math"
)

func main() {
	var input int
	fmt.Print("Enter number: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("please enter a valid number")
	}
	var inputPow float64
	fmt.Print("Enter pow: ")
	_, err = fmt.Scan(&inputPow)
	if err != nil {
		fmt.Println("please enter a valid pow")
	}
	result := math.Pow(float64(input), inputPow)
	fmt.Println("result:", result)
}
