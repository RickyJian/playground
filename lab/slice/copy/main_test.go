package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceCopy1(t *testing.T) {
	tests := []struct {
		new      []string
		old      []string
		expected []string
	}{
		{
			new:      []string{"1", "2", "3"},
			old:      []string{"4", "5", "6"},
			expected: []string{"1", "2", "3", "4", "5", "6"},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, SliceCopy1(test.new, test.old))
	}
}

func TestSliceCopy2(t *testing.T) {
	tests := []struct {
		new      []string
		old      []string
		expected []string
	}{
		{
			new:      []string{"1", "2", "3"},
			old:      []string{"4", "5", "6"},
			expected: []string{"1", "2", "3", "4", "5", "6"},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, SliceCopy2(test.new, test.old))
	}
}

func BenchmarkSliceCopy1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceCopy1(make([]string, i), make([]string, i))
	}
}

func BenchmarkSliceCopy2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceCopy2(make([]string, i), make([]string, i))
	}
}
