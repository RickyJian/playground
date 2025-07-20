package main

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
