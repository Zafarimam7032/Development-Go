package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	arr1 := [3]int{3, 4, 5}
	additionOfArray(arr, arr1)
	multiplicationOfArray(arr1, arr)
}

func additionOfArray(arr, arr1 [3]int) {
	var result [3]int
	for i := 0; i < 3; i++ {
		result[i] = arr[i] + arr1[i]
	}
	fmt.Println("addition:", result)

}

func multiplicationOfArray(arr, arr1 [3]int) {
	var result [3]int
	for i := 0; i < 3; i++ {
		result[i] = arr[i] * arr1[i]
	}
	fmt.Println("multiplication:", result)
}
