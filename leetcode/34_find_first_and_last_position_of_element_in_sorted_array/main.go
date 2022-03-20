package main

func main() {
}

func searchRangeBrute(nums []int, target int) []int {
	start, end := -1, -1
	for i, num := range nums {
		if num == target {
			if start == -1 {
				start = i
			}
			end = i
		}
	}
	return []int{start, end}
}

func searchRangeBinarySearch(nums []int, target int) []int {
	return []int{findFirstLeft(nums, target), findLastRight(nums, target)}
}

func findFirstLeft(nums []int, target int) int {
	result := -1
	startIndex, endIndex := 0, len(nums)-1
	for startIndex <= endIndex {
		index := (startIndex + endIndex) / 2
		if mid := nums[index]; mid == target {
			endIndex = index - 1
			result = index
		} else if mid < target {
			startIndex = index + 1
		} else {
			// mid > target
			endIndex = index - 1
		}
	}
	return result
}

func findLastRight(nums []int, target int) int {
	result := -1
	startIndex, endIndex := 0, len(nums)-1
	for startIndex <= endIndex {
		index := (startIndex + endIndex + 1) / 2
		if mid := nums[index]; mid == target {
			startIndex = index + 1
			result = index
		} else if mid < target {
			startIndex = index + 1
		} else {
			// mid > target
			endIndex = index - 1
		}
	}
	return result
}
