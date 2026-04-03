package main

import "fmt"

func main() {
	var number int
	var number1 int
	fmt.Print("Please Enter First Number:")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Print("Please Enter Valid Number:")
	}
	fmt.Print("Please Enter Second Number:")
	_, err1 := fmt.Scan(&number1)
	if err1 != nil {
		fmt.Print("Please Enter Valid Second Number:")
	}

	add := func(num int, num1 int) { fmt.Println("addition", num+num1) }
	mul := func(num int, num1 int) { fmt.Println("multiple", num*num1) }
	div := func(num int, num1 int) { fmt.Println("divide", num/num1) }
	sub := func(num int, num1 int) { fmt.Println("subtraction", num-num1) }

	calculate(number, number1, add)
	calculate(number, number1, mul)
	calculate(number, number1, div)
	calculate(number, number1, sub)
}

func calculate(num int, num1 int, operation func(num int, num1 int)) {
	operation(num, num1)
}
