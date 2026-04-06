package main

import "fmt"

func mapPractical() {
	fmt.Println("this is map practical function ")

	var m = map[int]int{1: 1}
	fmt.Println(m)
	fmt.Println(m[0])

	ma := make(map[int]bool)
	fmt.Println(ma)
	fmt.Println(ma[1])

}
