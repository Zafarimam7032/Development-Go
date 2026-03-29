package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("this is a main method")
	var num int8 = 127
	fmt.Printf("num = %d\n", num)

	var num1 = int16(num)
	fmt.Printf("num = %d", num1)

	var name = "zafar imam"
	fmt.Println(name)
	nameint, err := strconv.Atoi(name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("nameint = %d\n", nameint)

	name3, err := strconv.ParseBool("false")
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println(name3)
	}

	n, err := strconv.ParseInt("123456", 10, 32)
	fmt.Println("int", n)

	f, errr := strconv.ParseFloat("1.234", 8)
	if errr != nil {
	}
	fmt.Println("float", f)

	str := strconv.FormatInt(200, 10)
	fmt.Println(str)
}
