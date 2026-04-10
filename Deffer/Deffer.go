package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("First")
	defer fmt.Println("main defer")
	fmt.Println("Second")

	file, _ := os.Open("file.txt")
	defer file.Close()
	//defer work as finally in java
	defferFunction()
	PlayWithPanic()
	playWithPanic()
}

func defferFunction() {
	defer func() {
		if c := recover(); c != nil {
			fmt.Println(" panic : ", c)
		}
	}()
	fmt.Println("Deffer")
	panic("this is First panic")
}

func PlayWithPanic() {
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("without panic : ", c)
		}
	}()
}

func playWithPanic() {
	defer func() {
		if c := recover(); c != nil {
			fmt.Println(" with panic : ", c)
			defer func() {
				if rec := recover(); rec != nil {
					fmt.Println("inside")
				}
			}()
			panic("this is panic inside playWithPanic")
		}
	}()
	panic("this is more panic")
}
