package main

import (
	"fmt"
	"sort"
)

func main() {
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	results := make([][]int, 0)
	recursion(candidates, target, []int{}, &results)
	return results
}

func recursion(candidates []int, target int, currents []int, result *[][]int) {
	for i, candidate := range candidates {
		val := target - candidate
		if val < 0 {
			fmt.Println(target, candidates, currents, result)
			break
		} else if val == 0 {
			// append last candidate to new array and append to result
			curLen := len(currents)
			match := make([]int, curLen+1)
			copy(match, currents)
			match[curLen] = candidate
			// 重點
			*result = append(*result, match)
			break
		}

		// target > candidate
		recursion(candidates[i:], val, append(currents, candidate), result)
	}
}
