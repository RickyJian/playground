package main

import (
	"slices"
)

func main() {
	// TODO: implement here
}

func orangesRotting(grid [][]int) int {
	var starts [][2]int
	for row, cols := range grid {
		for col, cell := range cols {
			if cell == 2 {
				starts = append(starts, [2]int{row, col})
			}
		}
	}
	if len(starts) == 0 {
		for _, cols := range grid {
			if slices.Contains(cols, 1) {
				return -1
			}
		}
		return 0
	}
	step := bfs(starts, grid)
	for _, cols := range grid {
		if slices.Contains(cols, 1) {
			return -1
		}
	}
	return step
}

func bfs(starts [][2]int, grid [][]int) int {
	dirs := [][2]int{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}
	maxRow, maxCol := len(grid), len(grid[0])
	step := -1 // -1 is start posistion
	queue := starts
	for len(queue) > 0 {
		for _, first := range queue {
			queue = queue[1:]
			row, col := first[0], first[1]
			for _, dir := range dirs {
				nextRow, nextCol := row+dir[0], col+dir[1]
				if nextRow < 0 || nextCol < 0 || nextRow >= maxRow || nextCol >= maxCol {
					continue
				} else if cell := grid[nextRow][nextCol]; cell != 1 {
					continue
				}
				grid[nextRow][nextCol] = 0
				queue = append(queue, [2]int{nextRow, nextCol})
			}
		}
		step++
	}
	return step
}
