package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 測試案例
var testCases = []struct {
	name     string
	maze     [][]byte
	entrance []int
	expected int
}{
	{
		name: "範例 1",
		maze: [][]byte{
			{'+', '+', '.', '+'},
			{'.', '.', '.', '+'},
			{'+', '+', '+', '.'},
		},
		entrance: []int{1, 2},
		expected: 1,
	},
	{
		name: "範例 2",
		maze: [][]byte{
			{'+', '+', '+'},
			{'.', '.', '.'},
			{'+', '+', '+'},
		},
		entrance: []int{1, 0},
		expected: 2,
	},
	{
		name: "範例 3",
		maze: [][]byte{
			{'.', '+'},
		},
		entrance: []int{0, 0},
		expected: -1,
	},
	{
		name: "無法到達出口",
		maze: [][]byte{
			{'+', '+', '+'},
			{'+', '.', '+'},
			{'+', '+', '+'},
		},
		entrance: []int{1, 1},
		expected: -1,
	},
	{
		name: "入口即為出口",
		maze: [][]byte{
			{'.', '+'},
			{'+', '+'},
		},
		entrance: []int{0, 0},
		expected: -1,
	},
}

// 原始版本的測試
func TestNearestExit(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := nearestExit(tt.maze, tt.entrance)
			assert.Equal(t, tt.expected, result, "原始版本結果不符預期")
		})
	}
}

// 優化版本的測試
func TestNearestExitOptimized(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := nearestExitOptimized(tt.maze, tt.entrance)
			assert.Equal(t, tt.expected, result, "優化版本結果不符預期")
		})
	}
}

// 比較兩個版本的結果
func TestCompareResults(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result1 := nearestExit(tt.maze, tt.entrance)
			result2 := nearestExitOptimized(tt.maze, tt.entrance)
			assert.Equal(t, result1, result2, "原始版本和優化版本的結果不一致")
		})
	}
}

// 效能測試用的迷宮
var benchmarkMaze = [][]byte{
	{'+', '+', '.', '+', '+', '+', '+', '+', '+', '+'},
	{'+', '.', '.', '.', '.', '.', '.', '.', '.', '+'},
	{'+', '+', '+', '+', '+', '+', '+', '+', '.', '+'},
	{'+', '.', '.', '.', '.', '.', '.', '.', '.', '+'},
	{'+', '+', '+', '+', '+', '+', '+', '+', '.', '+'},
	{'+', '.', '.', '.', '.', '.', '.', '.', '.', '+'},
	{'+', '+', '+', '+', '+', '+', '+', '+', '.', '+'},
	{'+', '.', '.', '.', '.', '.', '.', '.', '.', '+'},
	{'+', '+', '+', '+', '+', '+', '+', '+', '.', '+'},
	{'+', '+', '+', '+', '+', '+', '+', '+', '+', '+'},
}
var benchmarkEntrance = []int{1, 1}

// 原始版本的效能測試
func BenchmarkNearestExit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nearestExit(benchmarkMaze, benchmarkEntrance)
	}
}

// 優化版本的效能測試
func BenchmarkNearestExitOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nearestExitOptimized(benchmarkMaze, benchmarkEntrance)
	}
}

// 優化版本 2 的測試
func TestNearestExitOptimized2(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// 複製迷宮以避免修改原始測試資料
			mazeCopy := make([][]byte, len(tt.maze))
			for i := range tt.maze {
				mazeCopy[i] = make([]byte, len(tt.maze[i]))
				copy(mazeCopy[i], tt.maze[i])
			}
			result := nearestExitOptimized2(mazeCopy, tt.entrance)
			assert.Equal(t, tt.expected, result, "優化版本 2 結果不符預期")
		})
	}
}

// 比較三個版本的結果
func TestCompareAllResults(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// 複製迷宮以避免修改原始測試資料
			mazeCopy1 := make([][]byte, len(tt.maze))
			mazeCopy2 := make([][]byte, len(tt.maze))
			mazeCopy3 := make([][]byte, len(tt.maze))
			for i := range tt.maze {
				mazeCopy1[i] = make([]byte, len(tt.maze[i]))
				mazeCopy2[i] = make([]byte, len(tt.maze[i]))
				mazeCopy3[i] = make([]byte, len(tt.maze[i]))
				copy(mazeCopy1[i], tt.maze[i])
				copy(mazeCopy2[i], tt.maze[i])
				copy(mazeCopy3[i], tt.maze[i])
			}
			result1 := nearestExit(mazeCopy1, tt.entrance)
			result2 := nearestExitOptimized(mazeCopy2, tt.entrance)
			result3 := nearestExitOptimized2(mazeCopy3, tt.entrance)
			assert.Equal(t, result1, result2, "原始版本和優化版本 1 的結果不一致")
			assert.Equal(t, result1, result3, "原始版本和優化版本 2 的結果不一致")
		})
	}
}

// 優化版本 2 的效能測試
func BenchmarkNearestExitOptimized2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 複製迷宮以避免修改原始測試資料
		mazeCopy := make([][]byte, len(benchmarkMaze))
		for i := range benchmarkMaze {
			mazeCopy[i] = make([]byte, len(benchmarkMaze[i]))
			copy(mazeCopy[i], benchmarkMaze[i])
		}
		nearestExitOptimized2(mazeCopy, benchmarkEntrance)
	}
}
