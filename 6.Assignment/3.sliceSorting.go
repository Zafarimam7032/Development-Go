package main

import (
	"fmt"
	"slices"
)

func main() {
	intSlice := []int{12, 90, 2, 79, 50, 60, 24, 12, 20, 10}

	slices.Sort(intSlice)
	fmt.Println("ascending Order", intSlice)
	slices.Reverse(intSlice)
	fmt.Println("Descending Order", intSlice)
}
