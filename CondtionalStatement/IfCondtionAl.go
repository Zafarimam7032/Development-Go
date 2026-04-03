package CondtionalStatement

import (
	"fmt"
)

func init() {
	fmt.Println("Init Condtional Statement")
}

func CondtionalStatement() {
	if 10 > 5 {
		fmt.Println("this will be true")
	}
	if 10 > 50 {
		fmt.Println("this will be true")
	} else {
		fmt.Println("this will be false")
	}
	if age := 10; age < 18 {
		fmt.Println("this will be true")
	}
}
