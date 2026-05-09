package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLastStoneWeight(t *testing.T) {
	tests := []struct {
		name     string
		stones   []int
		expected int
	}{
		{
			name:     "範例 1",
			stones:   []int{2, 7, 4, 1, 8, 1},
			expected: 1,
		},
		{
			name:     "範例 2",
			stones:   []int{1},
			expected: 1,
		},
		{
			name:     "全部可互相抵銷",
			stones:   []int{9, 3, 3, 3},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lastStoneWeight(tt.stones)
			assert.Equal(t, tt.expected, result)
		})
	}
}
