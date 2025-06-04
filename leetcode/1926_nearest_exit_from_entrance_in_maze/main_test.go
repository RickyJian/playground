package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNearestExit(t *testing.T) {
	tests := []struct {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := nearestExit(tt.maze, tt.entrance)
			assert.Equal(t, tt.expected, result)
		})
	}
}
