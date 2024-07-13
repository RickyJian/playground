package main

import (
	"sort"
)

func main() {
	// TODO: implement here
}

// Time Limit Exceeded
func maxOperationsV1(nums []int, k int) int {
	var count int
	used := make([]bool, len(nums))
	for i := 0; i < len(nums)-1; i++ {
		if used[i] {
			continue
		} else if nums[i] >= k {
			used[i] = true
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if used[j] {
				continue
			}
			if nums[i]+nums[j] == k {
				count++
				used[i], used[j] = true, true
				break
			}
		}
	}
	return count
}

// sorting solution
func maxOperationsV2(nums []int, k int) int {
	sort.Ints(nums)
	var count int
	for i, j := 0, len(nums)-1; i < j; {
		sum := nums[i] + nums[j]
		if nums[i]+nums[j] == k {
			count++
			i++
			j--
		} else if sum > k {
			j--
		} else {
			i++
		}
	}
	return count
}

// two pointer solution
func maxOperationsV3(nums []int, k int) int {
	// key is result, value is the same answer count
	resultMap := make(map[int]int)
	var count int
	for _, num := range nums {
		if _, ok := resultMap[num]; ok {
			count++
			resultMap[num]--
			if resultMap[num] == 0 {
				delete(resultMap, num)
			}
		} else {
			resultMap[k-num]++
		}
	}
	return count
}
