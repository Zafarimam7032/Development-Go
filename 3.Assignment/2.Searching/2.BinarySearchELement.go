package main

import (
	"fmt"
	"slices"
)

func Search(arr []int, search int) {
	i, flag := slices.BinarySearch(arr, search)
	if flag {
		fmt.Println("element found at index : ", i)
	}
}
