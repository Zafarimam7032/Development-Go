package main

import (
	"fmt"
)

func main() {
	var size int
	var slice []int
	fmt.Print("Enter Size:")
	_, err := fmt.Scan(&size)
	if err != nil {
		return
	}
	fmt.Println("Enter Elements:")
	for i := 0; i < size; i++ {
		var element int
		_, err2 := fmt.Scan(&element)
		if err2 != nil {
			return
		}
		slice = append(slice, element)
	}
	var search int
	fmt.Print("Enter The Search Elements:")
	_, err1 := fmt.Scan(&search)
	if err1 != nil {
		return
	}
	Sorting(slice)

	Search(slice, search)
}
