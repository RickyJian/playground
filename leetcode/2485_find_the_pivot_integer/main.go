package main

func main() {
	// TODO: implement here
}

// brute force
func pivotIntegerV1(n int) int {
	// init
	numbers := make([]int, n)
	for i := 1; i <= n; i++ {
		numbers[i-1] = i
	}
	// calculate
	for i := 0; i < n; i++ {
		left, right := numbers[:i+1], numbers[i:]
		var total int
		for _, val := range left {
			total += val
		}
		for _, val := range right {
			total -= val
		}
		if total == 0 {
			return i + 1
		}
	}
	return -1
}
