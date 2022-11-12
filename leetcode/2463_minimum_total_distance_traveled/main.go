package main

import (
	"math"
	"sort"
)

func main() {
	// TODO: implement here
}

// TLE
func minimumTotalDistanceDFS(robot []int, factory [][]int) int64 {
	sort.Ints(robot)
	sort.Slice(factory, func(i, j int) bool {
		return factory[i][0] < factory[j][0]
	})
	return dfs(robot, factory)
}

func dfs(robot []int, factory [][]int) int64 {
	if len(robot) == 0 || len(factory) == 0 {
		return 0
	}

	result, minPath := int64(math.MaxInt64), int64(math.MaxInt64)
	for _, f := range factory {
		if f[1] == 0 {
			continue
		}

		// robot[0]: robot position
		// f[0]: factory position
		d := abs(robot[0] - f[0])
		if d < minPath {
			f[1]--
			minPath = d
			current := dfs(robot[1:], factory) + d
			if current < result {
				result = current
			}
			f[1]++
		}
	}
	return result
}

func abs(number int) int64 {
	if number >= 0 {
		return int64(number)
	}
	return int64(-number)
}
