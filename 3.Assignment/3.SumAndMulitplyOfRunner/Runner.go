package main

import "fmt"

func main() {
	var sl []int
	var size int
	fmt.Println("Please enter the Size of Slices: ")
	_, err := fmt.Scan(&size)
	if err != nil {
		fmt.Println("please enter valid size ")
	}

	fmt.Println("please enter a elements")
	for i := 0; i < size; i++ {
		var element int
		_, err := fmt.Scan(&element)
		if err != nil {
			fmt.Println("please enter a valid element")
		}
		sl = append(sl, element)
	}
	AdditionOfAllElements(sl)
	MultiplicationOfAllElements(sl)
}
