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
	fmt.Println(maxProfit(nums), 5)
	fmt.Println(maxProfit(nums2), 1)
	fmt.Println(maxProfit(nums3), 1)
	fmt.Println(maxProfit(nums4), 2)
	fmt.Println(maxProfit(nums5), 5)
	fmt.Println(maxProfit(nums6), 6)
	fmt.Println(maxProfit(nums7), 2)
}

func maxProfit(prices []int) int {
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

func bestMaxProfit(prices []int) int {
	var max int
	buyPrice := prices[0]

	for _, price := range prices {
		if price < buyPrice {
			buyPrice = price
		}

		if max < price-buyPrice {
			max = price - buyPrice
		}
	}

	return max
}
