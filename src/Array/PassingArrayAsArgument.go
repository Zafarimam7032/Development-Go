package main

import "fmt"

func ArrayArgument() {
	arr := [3]int{1, 2, 3}
	fmt.Println("before passing as argument", arr)
	display(arr)
	fmt.Println("After passing as argument", arr)

	fmt.Println("before passing as Pointer argument", arr)
	show(&arr)
	fmt.Println("After passing as Pointer argument", arr)

}

func display(arr [3]int) {
	arr[0] = 20
	fmt.Println("after passing as argument arr =", arr)
}

func show(arr *[3]int) {
	arr[0] = 200
	fmt.Println("after passing as argument arr =", arr)
}
