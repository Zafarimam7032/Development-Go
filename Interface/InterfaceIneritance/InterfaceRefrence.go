package main

import "fmt"

type Interface1 interface {
	add(num int, num1 int)
}
type Interface3 interface {
	sub(num int, num1 int)
}

type Interface2 interface {
	Interface1
	Interface3
}
type Calculator struct {
}

func (c Calculator) add(num int, num1 int) {
	fmt.Println("Addition : ", num+num1)
}

func (c Calculator) sub(num int, num1 int) {
	fmt.Println("Substraction : ", num-num1)
}

func main() {
	cal := Calculator{}
	var in Interface2 = cal
	in.add(20, 30)
	var sub Interface3 = cal
	sub.sub(40, 10)
}
