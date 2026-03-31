package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Enter the Number:")
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Print("Please Enter Valid Number")
	}
	checkEvenOrOdd(number)
	checkWeatherNegativeOrPositive(number)
	checkCountOfDigit(number)
}

func checkEvenOrOdd(number int) {
	if number%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}
}

func checkWeatherNegativeOrPositive(number int) {
	if number >= 0 {
		fmt.Println("The number is positive")
	} else {
		fmt.Println("The number is negative")
	}
}

func checkCountOfDigit(number int) {
	//200
	count := 0
	positiveNumber := int32(math.Abs(float64(number)))
	for positiveNumber > 0 {
		count++
		positiveNumber = positiveNumber / 10
	}
	fmt.Println("Count of Digit", count)
}
