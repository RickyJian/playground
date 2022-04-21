package main

import (
	"math"
)

func main() {
}

func findClosestElementsS2E(arr []int, k int, x int) []int {
	for {
		if len(arr) == k {
			break
		}

		start, end := arr[0], arr[len(arr)-1]
		startClo := int(math.Abs(float64(x - start)))
		endClo := int(math.Abs(float64(x - end)))
		if startClo <= endClo {
			arr = arr[:len(arr)-1]
		} else {
			arr = arr[1:]
		}
	}
	return arr
}

// refer to: https://lenchen.medium.com/leetcode-658-find-k-closest-elements-1755777b9aaf
func findClosestElementsBS(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-k
	for left < right {
		midIdx := left + (right-left)/2
		// 若 x-arr[midIdx] > arr[midIdx+k]-x，
		// 代表 x 較靠近陣列右邊，arr[left:midIdx] 即可捨棄不看，
		// 反之，若 x 較靠近左邊，arr[midIdx+1:right] 即可捨棄不看。
		if x-arr[midIdx] > arr[midIdx+k]-x {
			left = midIdx + 1
		} else {
			right = midIdx
		}
	}
	return arr[left : left+k]
}
