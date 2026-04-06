package main

import "fmt"

type Person struct {
	id      int64
	name    string
	age     int
	address string
}

func (p Person) dis() Person {
	p.name = "Devid"
	fmt.Println("inside Method call : ", p)
	return p
}
