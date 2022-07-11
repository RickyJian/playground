package main

import (
	"sort"
)

func main() {
}

func combinationSum(candidates []int, target int) [][]int {
	// 由於題目未告知陣列的元素是否排序，因此在這先做排序待使用
	sort.Ints(candidates)
	// 回傳結果
	results := make([][]int, 0)
	recursion(candidates, target, []int{}, &results)
	return results
}

// recursion
//  * candidates：題目指定陣列元素
//  * target：題目指定合併結果
//  * currents：儲存符合元素，待後續寫回 result
//  * result：回傳結果
func recursion(candidates []int, target int, currents []int, result *[][]int) {
	for i, candidate := range candidates {
		val := target - candidate
		if val < 0 {
			// target < candidate
			// 由於傳進 function 中的陣列皆為升冪排序，因此當 val < 0 時代表後續元素皆無法匹配，
			// 因此直接跳出迴圈，並不會將結果回寫回 result，
			// 回到 recursion 上一層並走訪下一個元素
			break
		} else if val == 0 {
			// target == candidate
			// 當 target == candidate 代表當下 currents + candidate 可以組成的結果等於 target
			// 因此將 candidate 寫回 currents，再將 currents 寫回 result 跳出迴圈，
			// 回到 recursion 上一層並走訪下一個元素
			curLen := len(currents)
			match := make([]int, curLen+1)
			copy(match, currents)
			match[curLen] = candidate
			*result = append(*result, match)
			break
		}
		// target > candidate
		// 尚未滿足條件繼續將當下 candidate 寫進 currents，
		// 並繼續檢查當下組合是否在還有符合條件的元素
		recursion(candidates[i:], val, append(currents, candidate), result)
	}
}
