package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("GO routing ")
	go display("Hello World1", 1*time.Second)
	go display("Hello World2", 1*time.Second)
	go display("Hello World3", 1*time.Second)
	go display("Hello World4", 1*time.Second)
	go display("Hello World5", 1*time.Second)

	time.Sleep(1 * time.Second)
	fmt.Println("GO routing ended")
	time.Sleep(5 * time.Second)
}

func display(sentence string, waittime time.Duration) {
	time.Sleep(waittime)
	fmt.Println(sentence)
}
