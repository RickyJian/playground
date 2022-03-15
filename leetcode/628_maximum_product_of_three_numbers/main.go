package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// fmt.Println(maximumProduct([]int{1, 2, 3}))
	// fmt.Println(maximumProduct([]int{-1, -2, -3}))
	// fmt.Println(maximumProduct([]int{-100, -98, -1, 2, 3, 4}))
	// fmt.Println(maximumProduct([]int{-8, -7, -2, 10, 20}))
	// fmt.Println(maximumProductBrute([]int{-100, -2, -3, 1}))
	// fmt.Println(maximumProduct([]int{-100, -2, -3, 1}))
	// fmt.Println(maximumProductLoop([]int{-100, -2, -3, 1}))
	fmt.Println(maximumProductLoop([]int{1, 2, 3, 4}))
}

func maximumProductBrute(nums []int) int {
	result := math.MinInt64
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for z := j + 1; z < len(nums); z++ {
				if product := nums[i] * nums[j] * nums[z]; result < product {
					result = product
				}
			}
		}
	}
	return result
}

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	left1, left2 := nums[0], nums[1]
	right1, right2, right3 := nums[len(nums)-1], nums[len(nums)-2], nums[len(nums)-3]
	product1 := left1 * left2 * right1
	product2 := right1 * right2 * right3
	if product1 > product2 {
		return product1
	}
	return product2
}

func maximumProductLoop(nums []int) int {
	max1, max2, max3 := math.MinInt, math.MinInt, math.MinInt
	min1, min2 := math.MaxInt, math.MaxInt
	for _, num := range nums {
		if num >= max1 {
			max3, max2, max1 = max2, max1, num
		} else if num >= max2 {
			max3, max2 = max2, num
		} else if num >= max3 {
			max3 = num
		}
		if num <= min1 {
			min2, min1 = min1, num
		} else if num <= min2 {
			min2 = num
		}
	}
	product1, product2 := max1*max2*max3, min1*min2*max1
	if product1 > product2 {
		return product1
	}
	return product2
}
