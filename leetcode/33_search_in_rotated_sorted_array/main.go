package main

func main() {
}

func searchBrute(nums []int, target int) int {
	for i, num := range nums {
		if target == num {
			return i
		}
	}
	return -1
}

func searchBS(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		midIdx := (left + right) / 2
		mid := nums[midIdx]
		if target == mid {
			return midIdx
		} else if mid < nums[right] {
			// 右邊有序
			if mid < target && target <= nums[right] {
				// target 存在右邊有序數組，可直接做 Binary Search 搜尋
				left = midIdx + 1
			} else {
				// target 不存在於右邊有序數組，須向左邊數組做查詢
				right = midIdx - 1
			}
		} else {
			// mid > nums[right]
			// 左邊有序
			if nums[left] <= target && target < mid {
				// target 存在左邊有序數組，可直接做 Binary Search 搜尋
				right = midIdx - 1
			} else {
				// target 不存在於左邊有序數組，須向右邊數組做查詢
				left = midIdx + 1
			}
		}
	}
	return -1
}
