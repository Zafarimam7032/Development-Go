package main

import (
	"fmt"
	"slices"
)

func BasicSlice() {
	var basics []int

	fmt.Println(basics)

	usingMakeSlice := make([]int, 1, 6)
	fmt.Println(usingMakeSlice)
	usingMakeSlice[0] = 1
	fmt.Println(usingMakeSlice)
	usingMakeSlice = append(usingMakeSlice, 4)
	usingMakeSlice = append(usingMakeSlice, 9)
	usingMakeSlice = append(usingMakeSlice, 5)
	usingMakeSlice = append(usingMakeSlice, 11)
	usingMakeSlice = append(usingMakeSlice, 3)
	usingMakeSlice = append(usingMakeSlice, 1)
	fmt.Println(usingMakeSlice)
	fmt.Println(len(usingMakeSlice))
	fmt.Print(cap(usingMakeSlice))

	fmt.Println("before sorting slice", usingMakeSlice)
	slices.Sort(usingMakeSlice)
	fmt.Println("after sorting slice", usingMakeSlice)
	slices.Reverse(usingMakeSlice)
	fmt.Println("after reversed slice", usingMakeSlice)
	slices.All(usingMakeSlice)

	destSlice := make([]int, len(usingMakeSlice))
	copy(destSlice, usingMakeSlice)
	fmt.Println(destSlice)
	fmt.Println("destination slices ", destSlice)
	cloneslices := slices.Clone(destSlice)
	fmt.Println("destination slices : ", cloneslices)
	destSlice = slices.Insert(destSlice, 2, 8, 10)
	slices.Sort(destSlice)
	fmt.Println(usingMakeSlice)
	fmt.Println("slices : ", destSlice)
	num, check := slices.BinarySearch(destSlice, 11)
	fmt.Println(num)
	fmt.Println(check)

	slices.Delete(destSlice, 0, 3)
	fmt.Println(destSlice)
	fmt.Println("after completing deletion : ", usingMakeSlice)
	passingAsSlice(usingMakeSlice)
	fmt.Println("Printing after calling the Function :  ", usingMakeSlice)

	for i, v := range usingMakeSlice {
		fmt.Println("index : ", i, " Value : ", v)
	}

}

func passingAsSlice(sl []int) {
	fmt.Println("before modifying the slice : ", sl)
	sl[0] = 200
	fmt.Println("after modifying : ", sl)
}

//pass by reference only
