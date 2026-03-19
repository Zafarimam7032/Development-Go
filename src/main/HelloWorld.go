package main

import "fmt"
import "example.com/src/packageFolder"
import "github.com/src/variables"

func main() {
	fmt.Println("Hello World")
	packageFolder.Demo()
	variables.Variables()
	fmt.Println(variables.Language)
	//fmt.Println(variables.language)// access lebel restricted
	// packageFolder.show()
}
