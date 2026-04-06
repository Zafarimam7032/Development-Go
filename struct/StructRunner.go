package main

import "fmt"

func main() {
	per := Person{id: 1, name: "zafar Imam", age: 28, address: "pune"}
	fmt.Println(per)
	fmt.Println(per.address)
	fmt.Println(per.age)
	fmt.Println(per.id)
	fmt.Println(per.name)
	p := Person{1, "zafar Imam", 12, "pune"}
	fmt.Println(p)
	p.Show()
	fmt.Println("before : ", p)
	obj := p.dis()
	fmt.Println("After : ", obj)
	fmt.Println("after : ", p)

}

func (p Person) Show() {
	fmt.Println("Person ", p)
}
