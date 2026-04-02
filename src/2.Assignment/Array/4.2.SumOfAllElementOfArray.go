package main

import "fmt"

func SumOfAllElements(arr [4]int) {
	var result int
	for i := 0; i < len(arr); i++ {
		result = result + arr[i]
	}
	fmt.Println("Sum of all elements:", result)
}
