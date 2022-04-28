package main

import (
	"math"
	"sort"
)

func main() {
}

func smallestDistancePairBrute(nums []int, k int) int {
	results := make([]int, 0, 1000000)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			results = append(results, int(math.Abs(float64(nums[j]-nums[i]))))
		}
	}
	sort.Ints(results)
	return results[k-1]
}

func smallestDistancePairBrute2(nums []int, k int) int {
	counts := make([]int, 1000000)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			result := int(math.Abs(float64(nums[j] - nums[i])))
			counts[result]++
		}
	}
	var temp int
	for idx, count := range counts {
		temp += count
		if temp >= k {
			return idx
		}
	}
	return -1
}

func smallestDistancePairBS(nums []int, k int) int {
	sort.Ints(nums)

	// 假定有一個升冪排序的 `虛擬陣列`，
	// 陣列內元素皆為 nums[j] - nums[i] 的數值，
	// 可以得知此陣列的範圍介於：
	//   0 ~ nums[len(nums)-1]-nums[0]
	// 我們可以藉由
	//   * binary search：找出 `虛擬陣列` 中第 k 差值
	//   * slide window：統計差值 < 中間值的數量，來幫助 binary search 定位
	low, high := 0, nums[len(nums)-1]-nums[0]
	for low < high {
		mid := low + (high-low)/2
		var count, left int
		for right, num := range nums {
			for num-nums[left] > mid {
				left++
			}
			// 統計差值(nums[right] - nums[left]) < mid 的數量並與 k 做比對
			count += right - left
		}
		// 若數量 < k：中間值太小，須找出更大的中間值，因此將 low 改為 mid + 1
		//        > k：中間值太大，須找出更小的中間值，因此將 high 改為 mid
		//        = k：即是答案
		if count < k {
			low = mid + 1
		} else {
			// count >= k
			high = mid
		}
	}
	return low
}
