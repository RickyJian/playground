package main

import (
	"math"
)

func main() {
	// do nothing
}

// Time Limit Exceeded
func findKthLargest(nums []int, k int) int {
	pq := initQueue(len(nums))
	for _, num := range nums {
		pq.push(num)
	}
	return pq.findKth(k)
}

type priorityQueue []int

func initQueue(size int) priorityQueue {
	return make(priorityQueue, 0, size)
}

func (q *priorityQueue) push(num int) {
	// best way to use bs find position
	target := -1
	for i, element := range *q {
		if num > element {
			target = i
			break
		}
	}
	if target == -1 {
		*q = append(*q, num)
	} else {
		before := (*q)[:target]
		after := (*q)[target:]
		*q = make(priorityQueue, 0, cap((*q))+1)
		*q = append(*q, before...)
		*q = append(*q, num)
		*q = append(*q, after...)
	}
}

func (q *priorityQueue) findKth(k int) int {
	return (*q)[k-1]
}

// Time Limit Exceeded
func findKthLargestV2(nums []int, k int) int {
	pq := initQueueV2()
	for _, num := range nums {
		pq.push(num)
	}
	return pq.findKth(k)
}

type priorityQueueV2 struct {
	queue []int
}

func initQueueV2() *priorityQueueV2 {
	return &priorityQueueV2{}
}

func (q *priorityQueueV2) push(num int) {
	left, right := 0, len(q.queue)
	var position int
	for left < right {
		mid := (left + right) / 2
		if q.queue[mid] == num {
			for mid < right && q.queue[mid] == num {
				mid++
			}
			position = mid
			break
		} else if q.queue[mid] < num {
			// 向右插入
			left = mid + 1
		} else {
			// 向左插入
			right = mid
		}
	}
	// 如果沒有找到相同元素，left 就是插入位置
	if left == right {
		position = left
	}

	newQueue := make([]int, len(q.queue)+1)
	copy(newQueue, q.queue[:position])
	newQueue[position] = num
	copy(newQueue[position+1:], q.queue[position:])
	q.queue = newQueue
}

func (q *priorityQueueV2) findKth(k int) int {
	return q.queue[len(q.queue)-k]
}

// Time Limit Exceeded
func findKthLargestV3(nums []int, k int) int {
	pq := initQueueV3(k)
	for _, num := range nums {
		pq.push(num)
	}
	return pq[k-1]
}

type priorityQueueV3 []int

func initQueueV3(size int) priorityQueueV3 {
	return make(priorityQueueV3, 0, size)
}

func (p *priorityQueueV3) push(num int) {
	if size := len((*p)); size < cap((*p)) {
		(*p) = append((*p), num)
	} else if (*p)[size-1] < num {
		(*p) = (*p)[:size-1]
		(*p) = append((*p), num)
	} else {
		return
	}
	for i := len((*p)) - 1; i > 0 && (*p)[i] > (*p)[i-1]; i-- {
		(*p)[i], (*p)[i-1] = (*p)[i-1], (*p)[i]
	}
}

func findKthLargestV4(nums []int, k int) int {
	heap := newMaxHeap()
	for _, num := range nums {
		heap.push(num)
	}
	var last int
	for range k {
		last = heap.pop()
	}
	return last
}

type maxHeap struct {
	nums []int
}

func newMaxHeap() *maxHeap {
	return &maxHeap{}
}

func (h *maxHeap) pop() int {
	top := h.nums[0]
	h.nums[0] = h.nums[len(h.nums)-1]
	h.nums = h.nums[:len(h.nums)-1]
	if len(h.nums) == 0 {
		return top
	}

	// bubble down
	var currentIdx int
	for {
		var biggestIdx int
		biggestNode := math.MinInt
		size := len(h.nums)
		if leftIdx := 2*currentIdx + 1; leftIdx < size {
			if leftNode := h.nums[leftIdx]; leftNode > biggestNode {
				biggestNode = leftNode
				biggestIdx = leftIdx
			}
		}
		if rightIdx := 2*currentIdx + 2; rightIdx < size {
			if rightNode := h.nums[rightIdx]; rightNode > biggestNode {
				biggestNode = rightNode
				biggestIdx = rightIdx
			}
		}
		if currentNode := h.nums[currentIdx]; currentNode < biggestNode {
			h.nums[currentIdx], h.nums[biggestIdx] = h.nums[biggestIdx], h.nums[currentIdx]
			currentIdx = biggestIdx
		} else {
			break
		}
	}
	return top
}

func (h *maxHeap) push(num int) {
	h.nums = append(h.nums, num)
	currentIdx := len(h.nums) - 1
	for {
		currentNode := h.nums[currentIdx]
		parentIdx := (currentIdx - 1) / 2
		parentNode := h.nums[parentIdx]
		if currentNode > parentNode {
			h.nums[currentIdx], h.nums[parentIdx] = h.nums[parentIdx], h.nums[currentIdx]
			currentIdx = parentIdx
		} else {
			break
		}
	}
}

func findKthLargestV5(nums []int, k int) int {
	heap := newMinHeap(k)
	for _, num := range nums {
		heap.push(num)
	}
	return heap.nums[0]
}

type minHeap struct {
	nums  []int
	limit int
}

func newMinHeap(limit int) *minHeap {
	return &minHeap{limit: limit}
}

func (h *minHeap) push(num int) {
	if h.limit == len(h.nums) {
		if top := h.nums[0]; top >= num {
			return
		} else {
			// pop then bubble down
			h.nums[0] = num
			var currentIdx int
			for {
				var smallestIdx int
				smallestNode := math.MaxInt
				size := len(h.nums)
				if leftIdx := 2*currentIdx + 1; leftIdx < size {
					if leftNode := h.nums[leftIdx]; leftNode < smallestNode {
						smallestNode = leftNode
						smallestIdx = leftIdx
					}
				}
				if rightIdx := 2*currentIdx + 2; rightIdx < size {
					if rightNode := h.nums[rightIdx]; rightNode < smallestNode {
						smallestNode = rightNode
						smallestIdx = rightIdx
					}
				}
				if currentNode := h.nums[currentIdx]; currentNode > smallestNode {
					h.nums[currentIdx], h.nums[smallestIdx] = h.nums[smallestIdx], h.nums[currentIdx]
					currentIdx = smallestIdx
				} else {
					break
				}
			}
		}
	} else {
		// push then bubble up
		h.nums = append(h.nums, num)
		currentIdx := len(h.nums) - 1
		for {
			currentNode := h.nums[currentIdx]
			parentIdx := (currentIdx - 1) / 2
			parentNode := h.nums[parentIdx]
			if currentNode < parentNode {
				h.nums[currentIdx], h.nums[parentIdx] = h.nums[parentIdx], h.nums[currentIdx]
				currentIdx = parentIdx
			} else {
				break
			}
		}
	}
}
