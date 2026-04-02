package main

import (
	"fmt"
)

func main() {
	const size = 4
	var arr [4]int
	fmt.Println("Enter elements of array: ")
	var element int
	for i := 0; i < size; i++ {
	X:
		_, err := fmt.Scan(&element)
		if err != nil {
			fmt.Println("please enter valid element.", err)
			goto X
		}
		arr[i] = element
	}
	fmt.Println("Original array:", arr)
	ReverseArray(arr)
	SumOfAllElements(arr)
}
