package main

import (
	"fmt"
	"strings"
)

func main() {
	user := []string{"Mathew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aron", "Elizabeth"}

	coins := 50
	m := make(map[string]int)
	m["a"] = 1
	m["e"] = 1
	m["i"] = 2
	m["o"] = 3
	m["u"] = 4

	mapResult := make(map[string]int)
	for _, name := range user {
		countVowel := vowelOccurence(name)
		var coinsCount int
		for k, v := range countVowel {
			count := m[k]
			coinsCount += count * v

			if coinsCount >= coins {
				break
			}
			coins = coins - coinsCount

		}

		mapResult[name] = coinsCount
		if coins <= 2 {
			break
		}
	}
	fmt.Println(mapResult)
	fmt.Println(coins)
}

func vowelOccurence(name string) map[string]int {
	m := make(map[string]int)
	for _, ch := range name {
		if strings.EqualFold(string(ch), "a") || strings.EqualFold(string(ch), "e") || strings.EqualFold(string(ch), "i") || strings.EqualFold(string(ch), "u") || strings.EqualFold(string(ch), "o") {
			m[string(ch)]++
		}
	}
	fmt.Println("name : ", name, " occurrence : ", m)
	return m
}
