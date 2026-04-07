package main

import "fmt"

func Even(arr []int) {
	for _, v := range arr {
		if v%2 == 0 {
			fmt.Println(v)
		}
	}
}

func Odd(arr []int) {
	for _, v := range arr {
		if v%2 != 0 {
			fmt.Println(v)
		}
	}
}
