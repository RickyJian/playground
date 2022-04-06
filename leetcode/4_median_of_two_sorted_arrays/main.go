package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(findMedianSortedArraysLogMN([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArraysLogMN([]int{1, 3, 5}, []int{2, 4, 6}))
	fmt.Println(findMedianSortedArraysLogMN([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArraysLogMN([]int{2, 4, 6, 15}, []int{1, 7, 8, 10, 17}))
	fmt.Println(findMedianSortedArraysLogMN([]int{1, 4, 5, 7}, []int{2, 3, 8, 10, 1}))
	fmt.Println(findMedianSortedArraysLogMN([]int{1, 4, 5, 7, 9}, []int{2, 3, 8, 10, 1}))

}

func findMedianSortedArraysBrute(nums1 []int, nums2 []int) float64 {
	numSum := make([]int, 0, len(nums1)+len(nums2))
	numSum = append(numSum, nums1...)
	numSum = append(numSum, nums2...)
	sort.Ints(numSum)

	numLen := len(numSum)
	mediaIdx, remaining := numLen/2, math.Mod(float64(numLen), 2)
	var results float64
	if remaining == 0 {
		// even
		results = float64(numSum[mediaIdx-1]+numSum[mediaIdx]) / 2
	} else {
		// odd
		results = float64(numSum[mediaIdx])
	}
	return results
}

func findMedianSortedArraysLogMN(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		// 保證下方迴圈不會超出邊界
		nums1, nums2 = nums2, nums1
	}

	num1Len, num2Len := len(nums1), len(nums2)
	totalLeft := (num1Len + num2Len + 1) / 2
	left, right := 0, num1Len
	for left < right {
		i := (left+right)/2 + 1
		j := totalLeft - i
		if nums1[i-1] > nums2[j] {
			// 太靠右往左移動，下一輪搜索區間 [left, i - 1]
			right = i - 1
		} else {
			// 下一輪搜索區間 [i, right]
			left = i
		}
	}

	// left 是最符合的中間分隔線
	i := left
	var nums1Left float64
	if i == 0 {
		nums1Left = float64(math.MinInt)
	} else {
		nums1Left = float64(nums1[i-1])
	}
	var nums1Right float64
	if i == num1Len {
		nums1Right = float64(math.MaxInt)
	} else {
		nums1Right = float64(nums1[i])
	}
	j := totalLeft - i
	var nums2Left float64
	if j == 0 {
		nums2Left = float64(math.MinInt)
	} else {
		nums2Left = float64(nums2[j-1])
	}
	var nums2Right float64
	if j == num2Len {
		nums2Right = float64(math.MaxInt)
	} else {
		nums2Right = float64(nums2[j])
	}
	maxLeft := math.Max(nums1Left, nums2Left)
	if math.Mod(float64(num1Len+num2Len), 2) == 0 {
		minRight := math.Min(nums1Right, nums2Right)
		return (maxLeft + minRight) / 2
	}
	return maxLeft
}
