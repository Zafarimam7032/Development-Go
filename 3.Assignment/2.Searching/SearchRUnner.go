package main

import "fmt"

func main() {
	var size int
	var slice []int
	fmt.Print("Enter Size:")
	fmt.Scan(&size)
	fmt.Println("Enter Elements:")
	for i := 0; i < size; i++ {
		var element int
		fmt.Scan(&element)
		slice = append(slice, element)
	}
	var search int
	fmt.Print("Enter The Search Elements:")
	fmt.Scan(&search)
	Search(slice, search)
}
