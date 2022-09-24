package main

import (
	"fmt"
)

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func main() {
	// MinStack minStack = new MinStack();
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println("GetMin:", minStack.GetMin())
	minStack.Pop()
	fmt.Println("Top:", minStack.Top())
	fmt.Println("GetMin:", minStack.GetMin())
}

type MinStack struct {
	arr []int
	min []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.min) == 0 || val <= this.GetMin() {
		this.min = append(this.min, val)
	}
	this.arr = append(this.arr, val)
}

func (this *MinStack) Pop() {
	if this.arr[len(this.arr)-1] == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}
	this.arr = this.arr[:len(this.arr)-1]
}

func (this *MinStack) Top() int {
	return this.arr[len(this.arr)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
