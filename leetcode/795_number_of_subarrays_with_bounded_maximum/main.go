package main

func main() {
	// TODO: implement here
}

func numSubarrayBoundedMaxBruteForce(nums []int, left int, right int) int {
	var result int
	for i := 0; i < len(nums); i++ {
		l := nums[i]
		// 當最左邊的元素 > 右邊邊界
		if l > right {
			continue
		}

		// 子陣列最大元素
		max := l
		for j := i; j < len(nums); j++ {
			r := nums[j]
			// 若當下元素 > 右邊邊界則不須走訪後續元素
			if r > right {
				break
			}

			// 若當下元素 > max 將該元素替換成 max
			if max < r {
				max = r
			}
			// 當子陣列最大元素 >= 左方邊界滿足題目要求結果須 + 1
			if max >= left {
				result++
			}
		}
	}
	return result
}

func numSubarrayBoundedMaxDP(nums []int, left int, right int) int {
	var result, count int
	// -1：由於 nums index 是由 0 開始計算
	breakIdx := -1
	for i, num := range nums {
		if num > right {
			// TODO: description
			breakIdx = i
			count = 0
		} else if num < left {
			// 若當下 num < left 代表 subarray count 等於先前的數量
			result += count
		} else {
			// left <= num <= right
			// TODO: description
			count = i - breakIdx
			result += count
		}
	}
	return result
}
