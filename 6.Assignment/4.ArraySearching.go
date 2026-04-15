package main

import "fmt"

func main() {
	const size = 10
	var array [size]int
	fmt.Println("enter Element :")
	for i := 0; i < size; i++ {
		var element int
		fmt.Scan(&element)
		array[i] = element
	}
	fmt.Println("array", array)
	var searchNumber int
	fmt.Println("enter Search Number :")
	fmt.Scan(&searchNumber)
	search(array, searchNumber)
}
func search(arr [10]int, searchNumber int) {
	bool := true
	for i := 0; i < len(arr); i++ {
		if arr[i] == searchNumber {
			fmt.Println("Found Element : ", searchNumber)
			bool = false
		}
	}
	if bool {
		fmt.Println("Not Found :", searchNumber)
	}
}
