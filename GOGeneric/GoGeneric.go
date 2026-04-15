package main

import "fmt"

func main() {
	fmt.Println("this is generic practice")
	fmt.Println(sum(10, 20))
	fmt.Println(sum(10.50, 20.50))
	fmt.Println(sum("zafar ", "imam"))
}

//	func sum(a, b int) int {
//		return a + b
//	}
type Oper interface {
	~int | ~float64 | ~string
}

func sum1[T Oper](a T, b T) T {
	return a + b

}

func sum[T int | float64 | string](a T, b T) T {
	return a + b

}
