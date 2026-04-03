package main

import "fmt"

type Status int

const (
	Pending Status = iota + 10
	Approved
	Rejected
)

func main() {
	var s Status = Approved
	enumDeclaration()
	fmt.Println(s)
}

func enumDeclaration() {
	const (
		sunday = iota + 2
		monday
		tuesday
		wednesday
		thursday
		friday
		saturday
	)
	fmt.Println(sunday)
}
