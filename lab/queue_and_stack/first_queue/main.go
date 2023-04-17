package main

import (
	"fmt"
)

func main() {
	q := newQueue()
	q.enQueue(1)
	q.enQueue(2)
	q.enQueue(3)
	q.enQueue(4)
	fmt.Println(q.front())
	q.deQueue()
	fmt.Println(q.front())
	q.deQueue()
	fmt.Println(q.front())
	q.deQueue()
	fmt.Println(q.front())
	q.deQueue()
	fmt.Println(q.front())
	q.deQueue()
	fmt.Println(q.front())
}

func newQueue() queue {
	return make(queue, 0)
}

type queue []int32

func (q *queue) enQueue(number int32) {
	*q = append(*q, number)
}

func (q *queue) deQueue() {
	if !q.isEmpty() {
		*q = (*q)[1:]
	}
}

func (q *queue) isEmpty() bool {
	return len(*q) == 0
}

func (q *queue) front() int32 {
	if q.isEmpty() {
		return -1
	}
	return (*q)[0]
}
