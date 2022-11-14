package main

import (
	"math"
)

func main() {
	// TODO: implement here
}

// TLE
func minimumTotalDistanceDFS(robot []int, factory [][]int) int64 {
	return dfs(robot, factory)
}

func dfs(robot []int, factory [][]int) int64 {
	if len(robot) == 0 {
		return 0
	}

	result := int64(math.MaxInt64)
	for _, f := range factory {
		robotPos, factoryPos, limit := robot[0], f[0], f[1]
		if limit > 0 {
			f[1]--
			current := dfs(robot[1:], factory) + abs(robotPos-factoryPos)
			if current < result {
				result = current
			}
			f[1]++
		}
	}
	return result
}

func minimumTotalDistanceDFSMemo(robot []int, factory [][]int) int64 {
	dp := make([][]int, len(robot))
	for i := 0; i < len(robot); i++ {
		dp[i] = make([]int, len(factory))
	}
	return dfsMemo(robot, factory, 0, 0, dp)
}

func dfsMemo(robot []int, factory [][]int, robotPos, factoryPos int, dp [][]int) int64 {
	if len(robot) == 0 {
		return 0
	}

	// TODO: dfs with memo
	return int64(dp[robotPos][factoryPos])
}

func abs(number int) int64 {
	if number >= 0 {
		return int64(number)
	}
	return int64(-number)
}
