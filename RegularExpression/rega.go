package main

import (
	"fmt"
	"regexp"
)

func main() {
	exp := "\\d+"
	compile, err := regexp.Compile(exp)
	if err != nil {
		return
	}
	word := "java12 object13 word"

	allString := compile.FindAllString(word, -1)

	fmt.Println(allString)
}
