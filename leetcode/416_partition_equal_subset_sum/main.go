package main

import (
	"fmt"
)

func main() {
	fmt.Println(canPartitionDFS([]int{1, 5, 11, 5}))
	fmt.Println(canPartitionDFS([]int{1, 2, 3, 5}))
	fmt.Println(canPartitionDFS([]int{2, 2, 1, 1}))
	fmt.Println(canPartitionDFS([]int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
		100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
		100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
		100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
		100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
		100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
		100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
		100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
		100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
		100, 100, 100, 100, 100, 100, 99, 97}))
}

func canPartitionDFS(nums []int) bool {
	var total int
	for _, num := range nums {
		total += num
	}
	if total%2 > 0 {
		return false
	}

	half := total / 2
	for i := 0; i < len(nums); i++ {
		if dfs(nums, i, half, 0) {
			return true
		}
	}
	return false
}

func dfs(nums []int, idx int, target int, current int) bool {
	if target == current {
		return true
	} else if target < current {
		return false
	}

	for i := idx; i < len(nums); i++ {
		if dfs(nums, i+1, target, current+nums[i]) {
			return true
		}
	}
	return false
}
