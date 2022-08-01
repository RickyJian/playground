package main

import (
	"fmt"
)

func main() {
	q := newQueue(5)
	q.enQueue(1)
	q.enQueue(2)
	q.enQueue(3)
	q.enQueue(4)
	q.enQueue(5)
	q.deQueue()
	q.enQueue(6)
	fmt.Println(q.front(), q.rear())
}

func newQueue(size int) *queue {
	return &queue{
		arr:  make([]int32, size),
		head: -1,
		tail: -1,
	}
}

type queue struct {
	arr  []int32
	head int32
	tail int32
}

func (q *queue) enQueue(number int32) bool {
	if q.isFull() {
		return false
	} else if q.head == -1 {
		q.head = 0
	}

	q.tail = (q.tail + 1) % int32(len(q.arr))
	q.arr[q.tail] = number
	return true
}

func (q *queue) deQueue() bool {
	if q.isEmpty() {
		return false
	} else if q.head == q.tail {
		// when the head to the end, should reset head and tail
		q.head, q.tail = -1, -1
		return true
	}

	q.head = (q.head + 1) % int32(len(q.arr))
	return true
}

func (q *queue) isEmpty() bool {
	return q.head == -1
}

func (q *queue) isFull() bool {
	return q.head == (q.tail+1)%int32(len(q.arr))
}

func (q *queue) front() int32 {
	if q.isEmpty() {
		return -1
	}
	return q.arr[q.head]
}

func (q *queue) rear() int32 {
	if q.isEmpty() {
		return -1
	}
	return q.arr[q.tail]
}
