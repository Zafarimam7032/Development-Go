package main

import (
	"fmt"
	"sort"
)

func SortingArray() {
	arr := [4]int{6, 4, 5, 2}
	fmt.Println("before sorting array", arr)
	sort.Ints(arr[:])
	fmt.Println("after sorting array", arr)
	sort.Sort(sort.Reverse(sort.IntSlice(arr[:])))
	fmt.Println("after sorting array", arr)

	programmingLanguage := [4]string{"java", "python", "c", "golang"}
	fmt.Println("before sorting array", programmingLanguage)
	sort.Strings(programmingLanguage[:])
	fmt.Println("after sorting array", programmingLanguage)
	sort.Sort(sort.Reverse(sort.StringSlice(programmingLanguage[:])))
	fmt.Println("after sorting array in descending ", programmingLanguage)

}
