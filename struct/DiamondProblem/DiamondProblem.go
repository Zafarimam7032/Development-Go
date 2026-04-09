package main

import "fmt"

type A struct {
}

type B struct {
}

func (a A) hello() {
	fmt.Println("this is A Struct method")
}
func (b B) hello() {
	fmt.Println("this is B Struct method")
}

type C struct {
	A
	B
}

func main() {
	c := C{}
	c.A.hello()
	c.B.hello()
}
