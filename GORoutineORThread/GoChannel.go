package main

import (
	"fmt"
)

func main() {

	message := make(chan string)

	go func() {
		message <- "this is normal message"
		message <- "this is another message"
	}()

	fmt.Println(<-message)
	//close(message)
	fmt.Println(<-message)

	size := 5

	for i := 0; i < size; i++ {
		mess := fmt.Sprintf("this is hello World : %d", i)
		go sayHello(mess, message)
	}

	for i := 0; i < size; i++ {
		mess := <-message
		fmt.Println(mess)
	}

	num := make(chan int)
	go func() { num <- 1 }()
	fmt.Println(<-num)

}

func sayHello(message string, ch chan string) {
	ch <- message
}
