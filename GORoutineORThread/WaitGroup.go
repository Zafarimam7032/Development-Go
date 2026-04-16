package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(5)

	go SayHello("Hello World1", 1*time.Second, &wg)
	go SayHello("Hello World2", 1*time.Second, &wg)
	go SayHello("Hello World3", 1*time.Second, &wg)
	go SayHello("Hello World4", 1*time.Second, &wg)
	go SayHello("Hello World5", 1*time.Second, &wg)
	wg.Wait()
}

func SayHello(sentence string, wating time.Duration, wg *sync.WaitGroup) {
	time.Sleep(wating)
	fmt.Println(sentence)
	wg.Done()
}
