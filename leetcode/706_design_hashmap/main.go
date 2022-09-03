package main

import (
	"fmt"
)

func main() {
	m := Constructor()
	m.Put(1, 1)
	m.Put(2, 2)
	fmt.Println("1 Get:(1)", m.Get(1))
	fmt.Println("2 Get:(3)", m.Get(3))
	m.Put(2, 1)
	fmt.Println("3 Get:(2)", m.Get(2))
	m.Remove(2)
	fmt.Println("4 Get:(2)", m.Get(2))
}

type MyHashMap struct {
	firstNode *node
	node      *node
}

type node struct {
	key   int
	value int
	next  *node
}

func Constructor() MyHashMap {
	return MyHashMap{}
}

func (this *MyHashMap) Put(key int, value int) {
	firstNode := this.firstNode
	for {
		if firstNode == nil {
			this.firstNode = &node{
				key:   key,
				value: value,
			}
			this.node = this.firstNode
			break
		} else if key == firstNode.key {
			firstNode.value = value
			break
		} else if firstNode.next == nil {
			firstNode.next = &node{
				key:   key,
				value: value,
			}
			break
		}
		firstNode = firstNode.next
	}
}

func (this *MyHashMap) Get(key int) int {
	firstNode := this.firstNode
	for {
		if firstNode == nil {
			break
		} else if key == firstNode.key {
			return firstNode.value
		}
		firstNode = firstNode.next
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	firstNode := this.firstNode
	for {
		if firstNode == nil {
			break
		} else if key == firstNode.key {
			this.firstNode = firstNode.next
			this.node = firstNode.next
			break
		} else if firstNode.next != nil && key == firstNode.next.key {
			firstNode.next = firstNode.next.next
			break
		}
		firstNode = firstNode.next
	}
}

// TODO: slice
