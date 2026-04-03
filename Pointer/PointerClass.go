package main

import "fmt"

func main() {
	var a int = 10
	fmt.Println(a)
	var b *int = &a
	fmt.Println(b)

	name := "zafar imam"
	fmt.Println(name)

	pointerName := &name
	fmt.Println(pointerName)

}
