package main

import (
	"fmt"
	"regexp"
	"unicode"
)

func countWords(input string) int {
	re := regexp.MustCompile(`[A-Z][^A-Z]*`)
	words := re.FindAllString(input, -1)
	return len(words)
}

func main() {
	str := "saveChnageInThEditor"
	fmt.Println(countWords(str))
	countWord(str)
}

func countWord(input string) {
	count := 0
	word := ""

	for _, ch := range input {
		if unicode.IsUpper(ch) {
			if word != "" {
				count++
				word = ""
			}
		}
		word += string(ch)
	}

	if word != "" {
		count++
	}
	fmt.Println(count)
}
