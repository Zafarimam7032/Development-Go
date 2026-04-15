package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("1.Count of Character and word ")
	var sentence string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter text: ")
	scanner.Scan()
	sentence = scanner.Text()
	countOfChar := countOfCharacter(sentence)
	words := countOfWord(sentence)

	fmt.Println(words)
	fmt.Println(countOfChar)
}

func countOfCharacter(sentences string) map[string]int {
	strArr := strings.Split(sentences, "")
	result := make(map[string]int)
	for _, ch := range strArr {
		if ch != " " {
			result[ch]++
		}
	}
	return result
}

func countOfWord(sentences string) map[string]int {
	strArr := strings.Split(sentences, " ")
	result := make(map[string]int)
	for _, ch := range strArr {
		if ch != " " {
			result[ch]++
		}
	}
	return result
}
