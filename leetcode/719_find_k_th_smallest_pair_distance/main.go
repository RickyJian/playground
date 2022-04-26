package main

import (
	"math"
	"sort"
)

func main() {

}

func smallestDistancePairBrute(nums []int, k int) int {
	var results []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			results = append(results, int(math.Abs(float64(nums[j]-nums[i]))))
		}
	}
	sort.Ints(results)
	return results[k-1]
}
