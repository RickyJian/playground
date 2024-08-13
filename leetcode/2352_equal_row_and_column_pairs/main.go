package main

import (
	"strconv"
	"strings"
)

func main() {
	// TODO: implement here
}

func equalPairsV1(grid [][]int) int {
	rSet := make(map[string]int)
	cSet := make(map[string]int)
	for i := 0; i < len(grid); i++ {
		var rBuilder strings.Builder
		var cBuilder strings.Builder
		for j := 0; j < len(grid[i]); j++ {
			rBuilder.WriteString(strconv.Itoa(grid[i][j]))
			rBuilder.WriteString("_")
			cBuilder.WriteString(strconv.Itoa(grid[j][i]))
			cBuilder.WriteString("_")
		}
		rSet[rBuilder.String()]++
		cSet[cBuilder.String()]++
	}

	var result int
	for key, rCount := range rSet {
		val := rCount * cSet[key]
		result += val
	}
	return result
}
