package main

import (
	"fmt"
)

func main() {
	sl := make([]string, 5)
	sl[0] = "india"
	sl[1] = "canada"
	sl[2] = "Japan"
	sl[3] = "USA"
	sl[4] = "UK"
	sl = append(sl, "italy")
	sl = append(sl, "france")
	fmt.Println(sl)
	sl = removeElement(sl, "india")
	fmt.Println(sl)
}

func removeElement(slice []string, elem string) []string {
	var result []string
	for _, item := range slice {
		if item != elem {
			result = append(result, item)
		}
	}
	return result
}
