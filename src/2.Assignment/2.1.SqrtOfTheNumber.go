package main

import (
	"fmt"
	"math"
)

func main() {
	var number int
	fmt.Print("Please enter number: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Please enter valid number")
	}

	sqrt := math.Sqrt(float64(number))
	fmt.Println("Sqrt of number:", sqrt)

}
