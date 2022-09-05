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

	fmt.Println("======")

	m2 := newHashMapArr()
	m2.put(1, 1)
	m2.put(2, 2)
	fmt.Println("1 Get:(1)", m2.get(1))
	fmt.Println("2 Get:(3)", m2.get(3))
	m2.put(2, 1)
	fmt.Println("3 Get:(2)", m2.get(2))
	m2.remove(2)
	fmt.Println("4 Get:(2)", m2.get(2))

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

// worse case: Time Limit Exceeded
type hashMapArr struct {
	arr []*arr
}

type arr struct {
	key   int
	value int
}

func newHashMapArr() hashMapArr {
	return hashMapArr{arr: make([]*arr, 0, 1000001)}
}

func (h *hashMapArr) put(key int, value int) {
	for _, val := range h.arr {
		if val.key == key {
			val.value = value
			return
		}
	}
	h.arr = append(h.arr, &arr{key: key, value: value})
}

func (h *hashMapArr) get(key int) int {
	for _, val := range h.arr {
		if val.key == key {
			return val.value
		}
	}
	return -1
}

func (h *hashMapArr) remove(key int) {
	for i, val := range h.arr {
		if val.key == key {
			newArr := make([]*arr, 0, 10001)
			preArr, postArr := h.arr[:i], h.arr[i+1:]
			newArr = append(newArr, preArr...)
			newArr = append(newArr, postArr...)
			h.arr = newArr
			return
		}
	}
}
