package main

import "fmt"

func main() {
	normalFunction()
	mess := nomalFunctionWithReturnType()
	fmt.Println(mess)
	mess1, mess2 := normalFunctionWithMultipleReturnType()
	fmt.Println(mess1)
	fmt.Println(mess2)
}

func normalFunction() {
	fmt.Println("normalFunction")
}

func nomalFunctionWithReturnType() string {
	return "nomalFunctionWithReturnType"
}

func normalFunctionWithMultipleReturnType() (string, string) {

	return "nomalFunctionWithMultipleReturnType", "nomalFunctionWithMultipleReturnType"
}
