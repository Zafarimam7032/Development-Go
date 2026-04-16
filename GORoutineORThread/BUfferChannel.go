package main

import "fmt"

func main() {
	message := make(chan int, 2)
	message <- 1
	message <- 2
	go func() { message <- 3 }()
	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
}
