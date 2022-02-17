package main

import (
	"fmt"
)

func main() {
	fmt.Println(containsNearbyDuplicateV1([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicateV1([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicateV1([]int{1, 2, 3, 1, 2, 3}, 2))

	fmt.Println(containsNearbyDuplicateV2([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicateV2([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicateV2([]int{1, 2, 3, 1, 2, 3}, 2))

}

func containsNearbyDuplicateV1(nums []int, k int) bool {
	numMap := make(map[int]int, len(nums))
	for i, num := range nums {
		if position, ok := numMap[num]; !ok || i-position > k {
			numMap[num] = i
		} else if ok && i-position <= k {
			return true
		}
	}
	return false
}

func containsNearbyDuplicateV2(nums []int, k int) bool {
	numMap := make(map[int]int, len(nums))
	for i, num := range nums {
		if position, ok := numMap[num]; ok && i-position <= k {
			return true
		}
		numMap[num] = i
	}
	return false
}
