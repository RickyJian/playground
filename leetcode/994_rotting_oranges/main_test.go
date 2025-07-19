package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrangesRotting(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "範例 1",
			grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			expected: 4,
		},
		{
			name: "範例 2",
			grid: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			expected: -1,
		},
		{
			name: "範例 3",
			grid: [][]int{
				{0, 2},
			},
			expected: 0,
		},
		{
			name: "只有新鮮橘子",
			grid: [][]int{
				{1, 1, 1},
				{1, 1, 1},
			},
			expected: -1,
		},
		{
			name: "只有牆壁",
			grid: [][]int{
				{0},
			},
			expected: 0,
		},
		{
			name: "兩顆爛橘子",
			grid: [][]int{
				{0, 2, 2},
			},
			expected: 0,
		},
		{
			name: "兩顆爛橘子在對角線",
			grid: [][]int{
				{2, 1, 1},
				{1, 1, 1},
				{0, 1, 2},
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := orangesRotting(tt.grid)
			assert.Equal(t, tt.expected, result)
		})
	}
}
