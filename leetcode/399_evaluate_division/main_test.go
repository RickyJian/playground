package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcEquation(t *testing.T) {
	tests := []struct {
		name      string
		equations [][]string
		values    []float64
		queries   [][]string
		expected  []float64
	}{
		{
			name:      "Example 1",
			equations: [][]string{{"a", "b"}, {"b", "c"}},
			values:    []float64{2.0, 3.0},
			queries:   [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			expected:  []float64{6.0, 0.5, -1.0, 1.0, -1.0},
		},
		{
			name:      "Example 2",
			equations: [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
			values:    []float64{1.5, 2.5, 5.0},
			queries:   [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
			expected:  []float64{3.75, 0.4, 5.0, 0.2},
		},
		{
			name:      "Example 3",
			equations: [][]string{{"a", "b"}},
			values:    []float64{0.5},
			queries:   [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}},
			expected:  []float64{0.5, 2.0, -1.0, -1.0},
		},
		{
			name:      "Empty Input",
			equations: [][]string{},
			values:    []float64{},
			queries:   [][]string{{"a", "b"}},
			expected:  []float64{-1.0},
		},
		{
			name:      "Same Variable",
			equations: [][]string{{"a", "b"}},
			values:    []float64{2.0},
			queries:   [][]string{{"a", "a"}, {"b", "b"}},
			expected:  []float64{1.0, 1.0},
		},
	}
	for _, test := range tests {
		result := calcEquation(test.equations, test.values, test.queries)
		fmt.Println(result)
		assert.InDeltaSlice(t, test.expected, result, 0.00001)
	}
}
