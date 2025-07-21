package main

import (
	"fmt"
)

func main() {
	s := []int{2, 4, 6, 8, 10}
	fmt.Println(bsInsert(0, s))
	// bsInsert(1, s)
	fmt.Println(bsInsert(2, s))
	fmt.Println(bsInsert(7, s))
	fmt.Println(bsInsert(9, s))
}

func bsInsert(target int, s []int) []int {
	left, right := 0, len(s)
	var insertPos int
	for left < right {
		mid := (left + right) / 2
		if s[mid] == target {
			// find latest elements then append at the end of the sanme element
			for mid < right && s[mid] == target {
				mid++
			}
			insertPos = mid
			break
		} else if s[mid] < target {
			left = mid + 1
		} else {
			// s[mid] > target
			right = mid
		}
	}
	if left == right {
		insertPos = left
	}
	temp := make([]int, len(s)+1)
	copy(temp, s[:insertPos])
	temp[insertPos] = target
	copy(temp[insertPos+1:], s[insertPos:])
	return temp
}
