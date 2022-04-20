package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(2, findClosestElementsBS([]int{2, 4, 6}, 2, 2))
	fmt.Println(3, findClosestElementsBS([]int{2, 4, 6}, 2, 3))
	fmt.Println(4, findClosestElementsBS([]int{2, 4, 6}, 2, 4))
	fmt.Println(5, findClosestElementsBS([]int{2, 4, 6}, 2, 5))
	fmt.Println(6, findClosestElementsBS([]int{2, 4, 6}, 2, 6))

	// fmt.Println(2, findClosestElementsS2E([]int{2, 4, 6}, 2, 2))
	// fmt.Println(3, findClosestElementsS2E([]int{2, 4, 6}, 2, 3))
	// fmt.Println(4, findClosestElementsS2E([]int{2, 4, 6}, 2, 4))
	// fmt.Println(5, findClosestElementsS2E([]int{2, 4, 6}, 2, 5))
	// fmt.Println(6, findClosestElementsS2E([]int{2, 4, 6}, 2, 6))
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

func findClosestElementsBS(arr []int, k int, x int) []int {
	if x <= arr[0] {
		return arr[:k]
	} else if x >= arr[len(arr)-1] {
		return arr[len(arr)-k:]
	}

	var idx int
	left, right := 0, len(arr)-1
	for left <= right {
		midIdx := left + (right-left)/2
		midValue := arr[midIdx]
		if midValue == x {
			idx = midIdx
			break
		} else if midValue < x {
			left = midIdx + 1
			idx = left
		} else {
			// midVale > x
			right = midIdx - 1
			idx = right
		}
	}

	// TODO: fix and enhance
	var results []int
	leftIdx, rightIdx := idx, idx+1
	for i := 0; i < k; i++ {
		left, right := arr[leftIdx], arr[rightIdx]
		leftClo := int(math.Abs(float64(x - left)))
		rightClo := int(math.Abs(float64(x - right)))
		if leftClo <= rightClo {
			results = append([]int{left}, results...)
			leftIdx--
		} else {
			results = append(results, right)
			rightIdx++
		}
	}
	return results
}
