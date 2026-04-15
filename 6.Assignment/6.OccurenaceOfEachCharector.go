package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "i am reading go programming language"

	m, sl := occuranceOfEachCharector(sentence)

	for _, v := range sl {
		fmt.Println(v, " : ", m[v])
	}

}

func occuranceOfEachCharector(sentence string) (map[string]int, []string) {
	m := make(map[string]int)
	strArr := strings.Split(sentence, "")
	strsl := make([]string, 0)
	for i := 0; i < len(strArr); i++ {
		if strArr[i] != " " {
			if _, ok := m[strArr[i]]; !ok {
				strsl = append(strsl, strArr[i])
			}
			m[strArr[i]]++
		}
	}
	return m, strsl
}
