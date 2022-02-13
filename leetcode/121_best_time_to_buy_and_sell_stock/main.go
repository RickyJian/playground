package main

import (
	"fmt"
)

var (
	nums  = []int{7, 1, 5, 3, 6, 4}
	nums2 = []int{2, 1, 2, 0, 1}
	nums3 = []int{1, 2}
	nums4 = []int{2, 1, 2, 0, 1, 2}
	nums5 = []int{7, 1, 5, 3, 6, 4}
	nums6 = []int{6, 1, 3, 2, 4, 7}
	nums7 = []int{2, 1, 2, 1, 0, 1, 2}
)

func main() {

	fmt.Println(maxProfitV1(nums), 5)
	fmt.Println(maxProfitV1(nums2), 1)
	fmt.Println(maxProfitV1(nums3), 1)
	fmt.Println(maxProfitV1(nums4), 2)
	fmt.Println(maxProfitV1(nums5), 5)
	fmt.Println(maxProfitV1(nums6), 6)
	fmt.Println(maxProfitV1(nums7), 2)

	fmt.Println(maxProfitV2(nums), 5)
	fmt.Println(maxProfitV2(nums2), 1)
	fmt.Println(maxProfitV2(nums3), 1)
	fmt.Println(maxProfitV2(nums4), 2)
	fmt.Println(maxProfitV2(nums5), 5)
	fmt.Println(maxProfitV2(nums6), 6)
	fmt.Println(maxProfitV2(nums7), 2)
}

func maxProfitV1(prices []int) int {
	priceLen := len(prices)
	if priceLen <= 1 {
		return 0
	}

	var max int
	for leftIdx, rightIdx := 0, 1; rightIdx < priceLen; rightIdx++ {
		left, right := prices[leftIdx], prices[rightIdx]
		if profit := right - left; profit > max {
			max = profit
		}
		if left > right {
			leftIdx = rightIdx
		}
	}
	return max
}

func maxProfitV2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	var max, buy int
	for i, price := range prices {
		if i == 0 || price < buy {
			buy = price
		}
		if sell := price - buy; sell > max {
			max = sell
		}
	}
	return max
}
