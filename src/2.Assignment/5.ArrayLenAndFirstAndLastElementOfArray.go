package main

import "fmt"

func main() {
	const size int = 4
	var arr [4]int
	fmt.Println("Enter the elements of the array...")
	var element int
	for i := 0; i < size; i++ {
	X:
		_, err := fmt.Scan(&element)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			goto X
		}
		arr[i] = element
	}

	fmt.Println("The array is:", arr)

	fmt.Println("Length of the array:", len(arr))
	fmt.Println("First element:", arr[0])
	fmt.Println("Last element:", arr[len(arr)-1])

}
