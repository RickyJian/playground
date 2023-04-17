package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepend1(t *testing.T) {
	tests := []struct {
		no       int
		nos      []int
		expected []int
	}{
		{
			no:       1,
			nos:      []int{2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, Prepend1(test.no, test.nos))
	}
}

func TestPrepend2(t *testing.T) {
	tests := []struct {
		no       int
		nos      []int
		expected []int
	}{
		{
			no:       1,
			nos:      []int{2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, Prepend2(test.no, test.nos))
	}
}

func BenchmarkPrepend1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Prepend1(i, make([]int, i+1))
	}
}

func BenchmarkPrepend2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Prepend2(i, make([]int, i+1))
	}
}
