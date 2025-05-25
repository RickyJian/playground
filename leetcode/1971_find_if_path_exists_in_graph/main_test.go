package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPath(t *testing.T) {
	var tests = []*struct {
		name        string
		n           int
		edges       [][]int
		source      int
		destination int
		expected    bool
	}{
		{
			name:        "Example 1",
			n:           3,
			edges:       [][]int{{0, 1}, {1, 2}, {2, 0}},
			source:      0,
			destination: 2,
			expected:    true,
		},
		{
			name:        "Example 2",
			n:           6,
			edges:       [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}},
			source:      0,
			destination: 5,
			expected:    false,
		},
		{
			name:        "Single Node",
			n:           1,
			edges:       [][]int{},
			source:      0,
			destination: 0,
			expected:    true,
		},
		{
			name:        "No Path",
			n:           4,
			edges:       [][]int{{0, 1}, {2, 3}},
			source:      0,
			destination: 3,
			expected:    false,
		},
		{
			name:        "Path from 5 to 9",
			n:           10,
			edges:       [][]int{{4, 3}, {1, 4}, {4, 8}, {1, 7}, {6, 4}, {4, 2}, {7, 4}, {4, 0}, {0, 9}, {5, 4}},
			source:      5,
			destination: 9,
			expected:    true,
		},
		{
			name:        "Path from 7 to 5",
			n:           10,
			edges:       [][]int{{4, 3}, {1, 4}, {4, 8}, {1, 7}, {6, 4}, {4, 2}, {7, 4}, {4, 0}, {0, 9}, {5, 4}},
			source:      7,
			destination: 5,
			expected:    true,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, validPath(test.n, test.edges, test.source, test.destination), test.name)
	}
}

func TestValidPathLoopDFS(t *testing.T) {
	var tests = []*struct {
		name        string
		n           int
		edges       [][]int
		source      int
		destination int
		expected    bool
	}{
		{
			name:        "Example 1",
			n:           3,
			edges:       [][]int{{0, 1}, {1, 2}, {2, 0}},
			source:      0,
			destination: 2,
			expected:    true,
		},
		{
			name:        "Example 2",
			n:           6,
			edges:       [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}},
			source:      0,
			destination: 5,
			expected:    false,
		},
		{
			name:        "Single Node",
			n:           1,
			edges:       [][]int{},
			source:      0,
			destination: 0,
			expected:    true,
		},
		{
			name:        "No Path",
			n:           4,
			edges:       [][]int{{0, 1}, {2, 3}},
			source:      0,
			destination: 3,
			expected:    false,
		},
		{
			name:        "Path from 5 to 9",
			n:           10,
			edges:       [][]int{{4, 3}, {1, 4}, {4, 8}, {1, 7}, {6, 4}, {4, 2}, {7, 4}, {4, 0}, {0, 9}, {5, 4}},
			source:      5,
			destination: 9,
			expected:    true,
		},
		{
			name:        "Path from 7 to 5",
			n:           10,
			edges:       [][]int{{4, 3}, {1, 4}, {4, 8}, {1, 7}, {6, 4}, {4, 2}, {7, 4}, {4, 0}, {0, 9}, {5, 4}},
			source:      7,
			destination: 5,
			expected:    true,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, validPathLoopDFS(test.n, test.edges, test.source, test.destination), test.name)
	}
}

func TestValidPathBFS(t *testing.T) {
	var tests = []*struct {
		name        string
		n           int
		edges       [][]int
		source      int
		destination int
		expected    bool
	}{
		{
			name:        "Example 1",
			n:           3,
			edges:       [][]int{{0, 1}, {1, 2}, {2, 0}},
			source:      0,
			destination: 2,
			expected:    true,
		},
		{
			name:        "Example 2",
			n:           6,
			edges:       [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}},
			source:      0,
			destination: 5,
			expected:    false,
		},
		{
			name:        "Single Node",
			n:           1,
			edges:       [][]int{},
			source:      0,
			destination: 0,
			expected:    true,
		},
		{
			name:        "No Path",
			n:           4,
			edges:       [][]int{{0, 1}, {2, 3}},
			source:      0,
			destination: 3,
			expected:    false,
		},
		{
			name:        "Path from 5 to 9",
			n:           10,
			edges:       [][]int{{4, 3}, {1, 4}, {4, 8}, {1, 7}, {6, 4}, {4, 2}, {7, 4}, {4, 0}, {0, 9}, {5, 4}},
			source:      5,
			destination: 9,
			expected:    true,
		},
		{
			name:        "Path from 7 to 5",
			n:           10,
			edges:       [][]int{{4, 3}, {1, 4}, {4, 8}, {1, 7}, {6, 4}, {4, 2}, {7, 4}, {4, 0}, {0, 9}, {5, 4}},
			source:      7,
			destination: 5,
			expected:    true,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, validPathBFS(test.n, test.edges, test.source, test.destination), test.name)
	}
}
