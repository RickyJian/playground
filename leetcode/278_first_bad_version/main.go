package main

import (
	"fmt"
)

func main() {
	fmt.Println(firstBadVersion(1))
	fmt.Println(firstBadVersion(2))
	fmt.Println(firstBadVersion(3))
	fmt.Println(firstBadVersion(4))
	fmt.Println(firstBadVersion(5))
	fmt.Println(firstBadVersion(6))
	fmt.Println(firstBadVersion(6))
	fmt.Println(firstBadVersion(7))
}

func firstBadVersion(n int) int {
	index := -1
	start, end := 1, n
	for start <= end {
		n = (start + end) / 2
		if isBadVersion(n) {
			index = n
			end = n - 1
		} else {
			start = n + 1
		}
	}
	return index
}

var arr = map[int]bool{
	1: false,
	2: false,
	3: true,
	4: true,
	5: true,
	6: true,
	7: true,
}

func isBadVersion(version int) bool {
	return arr[version]
}
