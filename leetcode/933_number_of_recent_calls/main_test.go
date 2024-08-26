package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingV1(t *testing.T) {
	var tests = []*struct {
		reqs     []int
		expected []int
	}{
		{
			reqs:     []int{1, 100, 3001, 3002},
			expected: []int{1, 2, 3, 3},
		},
		{
			reqs:     []int{642, 1849, 4921, 5936, 5957},
			expected: []int{1, 2, 1, 2, 3},
		},
	}
	for _, test := range tests {
		obj := Constructor()
		for i, req := range test.reqs {
			assert.Equal(t, test.expected[i], obj.PingV1(req), req)
		}
	}
}
