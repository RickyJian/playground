package main

func main() {
	// do nothing
}

func lastStoneWeight(stones []int) int {
	h := newMaxHeap()
	for _, stone := range stones {
		h.push(stone)
	}
	for {
		if remaining := len(h.stones); remaining == 0 {
			return 0
		} else if remaining == 1 {
			return h.stones[0]
		}

		f, s := h.pop(), h.pop()
		if result := f - s; result > 0 {
			h.push(result)
		}
	}
}

type maxHeap struct {
	stones []int
}

func newMaxHeap() *maxHeap {
	return &maxHeap{}
}

func (h *maxHeap) pop() int {
	top := h.stones[0]
	h.stones[0] = h.stones[len(h.stones)-1]
	h.stones = h.stones[:len(h.stones)-1]
	size := len(h.stones)
	if size == 0 {
		return top
	}

	// bubble down
	var currentIdx int
	for {
		var biggest, childIdx int
		if left := 2*currentIdx + 1; left < size && h.stones[left] > biggest {
			biggest = h.stones[left]
			childIdx = left
		}
		if right := 2*currentIdx + 2; right < size && h.stones[right] > biggest {
			biggest = h.stones[right]
			childIdx = right
		}
		if h.stones[currentIdx] < biggest {
			h.stones[currentIdx], h.stones[childIdx] = h.stones[childIdx], h.stones[currentIdx]
			currentIdx = childIdx
		} else {
			return top
		}
	}
}

func (h *maxHeap) push(stone int) {
	h.stones = append(h.stones, stone)
	currentIdx := len(h.stones) - 1
	for {
		parentNodeIdx := (currentIdx - 1) / 2
		parentNode := h.stones[parentNodeIdx]
		currentNode := h.stones[currentIdx]
		if currentNode > parentNode {
			h.stones[parentNodeIdx], h.stones[currentIdx] = h.stones[currentIdx], h.stones[parentNodeIdx]
			currentIdx = parentNodeIdx
		} else {
			break
		}
	}
}
