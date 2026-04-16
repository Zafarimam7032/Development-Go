package main

import "fmt"

type QueOperation interface {
	insert(date int)
	remove()
	display()
}
type Queue struct {
	data []int
}

func (q *Queue) insert(date int) {
	q.data = append(q.data, date)
}
func (q *Queue) remove() {
	q.data = q.data[1:]
}
func (q *Queue) display() {
	fmt.Println(q.data)
}

func main() {
	queue := Queue{make([]int, 0)}
	queue.insert(1)
	queue.insert(2)
	queue.insert(3)
	queue.insert(4)
	queue.display()
	queue.remove()
	queue.display()
	queue.remove()
	queue.display()
}
