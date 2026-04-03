package main

import (
	"fmt"
	"math"
)

func main() {
	var intNum int = 10000
	fmt.Println(intNum)

	var floatNum float64 = 1.1
	fmt.Println(floatNum)
	var check bool = true
	fmt.Println(check)
	var num complex64 = complex64(1.1)
	fmt.Println(num)
	var str string = "this is a string"
	fmt.Println(str)
	fmt.Println(len(str))
	var char byte = 'a'
	fmt.Println(char)
	var ch rune = 'a'
	fmt.Println(string(ch))
	fmt.Println("min and max", math.MinInt8, math.MaxInt8)
	fmt.Println("min and max", math.MinInt16, math.MaxInt16)
	fmt.Println("min and max", math.MinInt32, math.MaxInt32)
	fmt.Println("min and max", math.MinInt64, math.MaxInt64)
	fmt.Println("min and max", math.MinInt, math.MaxInt)

	fmt.Println("float min and max", math.SmallestNonzeroFloat64, math.MaxFloat64)
	fmt.Println("float min and max", math.SmallestNonzeroFloat32, math.MaxFloat32)

}
