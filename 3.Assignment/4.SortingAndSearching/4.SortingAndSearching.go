package main

import (
	"fmt"
	"slices"
	"sort"
)

func Sorting(slice []int) {
	sort.Ints(slice)
	fmt.Println("after sorting", slice)
	_ = sort.Reverse(sort.IntSlice(slice))
	fmt.Println("after reverse", slice)
	slices.Sort(slice)
	fmt.Println("after sorting", slice)
	slices.Reverse(slice)
	fmt.Println("after reverse", slice)

}

func Search(slice []int, search int) {
	fmt.Println("searching", slice, search)
	slices.Sort(slice)
	i, flag := slices.BinarySearch(slice, search)
	if flag {
		fmt.Println("Search Element:", i)
	}
}
