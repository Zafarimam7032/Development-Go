package main

import "fmt"
import "example.com/src/packageFolder"
import "github.com/src/variables"
import "example.com/src/loop"
import "example.com/src/CondtionalStatement"

func main() {
	fmt.Println("Hello World")
	packageFolder.Demo()
	variables.Variables()
	loop.Loop()
	fmt.Println(variables.Language)
	//fmt.Println(variables.language)// access lebel restricted
	// packageFolder.show()
	CondtionalStatement.CondtionalStatement()
	CondtionalStatement.SwitchCase(1)
}
