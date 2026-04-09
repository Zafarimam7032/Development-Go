package main

import (
	"fmt"
	"reflect"
)

func main() {
	add(1, 2, 3, 4, 5, 6)
}

func add(numbers ...int) {
	fmt.Println(reflect.TypeOf(numbers).Kind()) //slice
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	fmt.Println(sum)
}
