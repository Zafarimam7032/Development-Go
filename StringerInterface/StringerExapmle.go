package main

import (
	"fmt"
	"strconv"
)

type MyInt int

func (num MyInt) String() string {

	return "Hello World"
}

type Employee struct {
	name string
	age  int
}

func (c Employee) String() string { // tostring like java

	return c.name + "" + strconv.Itoa(c.age)
}

type StringManupulation string

func (word StringManupulation) String() string {

	return "hi this is java Style string method" + string(word)
}

func main() {
	var num MyInt = 42
	fmt.Println(num)
	emp := Employee{name: "zafar", age: 24}
	fmt.Println(emp)
	var stringManupulation StringManupulation = "Zafar Imam"
	fmt.Println(stringManupulation)
}
