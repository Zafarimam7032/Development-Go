package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("this is main method")
	add, mult, div := multipleReturnType(10, 20)
	fmt.Println(add)
	fmt.Println(mult)
	fmt.Println(div)
	fmt.Println(reflect.TypeOf(div))

}

func multipleReturnType(num int, num1 int) (int, int, float32) {
	add := num + num1
	multi := num * num1
	div := num / num1

	return add, multi, float32(div)

}
