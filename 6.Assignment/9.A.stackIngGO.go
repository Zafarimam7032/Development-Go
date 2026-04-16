package main

import "fmt"

type Operation interface {
	insert(date int)
	remove()
	display()
}

type Stack struct {
	data []int
}

func (s *Stack) insert(date int) {
	s.data = append(s.data, date)
}
func (s *Stack) remove() {
	s.data = s.data[:len(s.data)-1]
}
func (s *Stack) display() {
	fmt.Println(s.data)
}

func main() {
	stack := Stack{data: make([]int, 0)}
	stack.insert(1)
	stack.insert(2)
	stack.insert(3)
	stack.display()
	stack.remove()
	stack.display()
}
