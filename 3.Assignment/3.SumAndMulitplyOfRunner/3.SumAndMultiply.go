package main

import "fmt"

func AdditionOfAllElements(slice []int) {
	var sum int
	for _, v := range slice {
		sum += v
	}
	fmt.Println("addition : ", sum)
}

func MultiplicationOfAllElements(slice []int) {
	var mul int = 1
	for _, v := range slice {
		mul *= v
	}
	fmt.Println("multiplication : ", mul)
}
