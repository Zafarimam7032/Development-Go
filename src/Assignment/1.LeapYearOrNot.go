package main

import (
	"fmt"
)

func main() {
	fmt.Print("Please Enter The Year:")
	var year int
X:
	_, err := fmt.Scan(&year)
	if err != nil {
		fmt.Print("Please Enter Valid Year:")
		goto X
	}

	if ValidLeapCheck(year) {
		fmt.Println("Leap Year")
	} else {
		fmt.Println("Not Leap Year")
	}

}

func ValidLeapCheck(year int) bool {
	if year%4 == 0 && year%100 != 0 {
		return true
	} else if year%400 == 0 {
		return true
	}
	return false
}
