package main

import "fmt"

func main() {
	show := func() string {
		return "hello"
	}

	mes := callUsingFunction(show)
	fmt.Println("message", mes())

	add := func(num int, num1 int) {
		fmt.Println("add:", +num+num1)
	}

	sub := func(num int, num1 int) {
		fmt.Println("sub:", +num-num1)
	}
	display(10, 20, add)
	display(20, 20, sub)

}

func display(num int, num1 int, operation func(int, int)) {
	operation(num, num1)
}

func callUsingFunction(fun func() string) func() string {
	return fun
}
