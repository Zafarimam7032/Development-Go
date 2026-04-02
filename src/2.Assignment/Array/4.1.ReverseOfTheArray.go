package main

import "fmt"

func ReverseArray(arr [4]int) {
	const size = len(arr)
	var reversedArr [size]int
	for i := size - 1; i >= 0; i-- {
		reversedArr[i] = arr[i]
	}
	fmt.Println("reversed Array", reversedArr)

}
