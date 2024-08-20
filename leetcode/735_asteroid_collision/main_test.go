package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsteroidCollision(t *testing.T) {
	var tests = []*struct {
		asteroids []int
		expected  []int
	}{
		{
			asteroids: []int{5, 10, -5},
			expected:  []int{5, 10},
		},
		{
			asteroids: []int{8, -8},
			expected:  []int{},
		},
		{
			asteroids: []int{10, 2, -5},
			expected:  []int{10},
		},
		{
			asteroids: []int{-2, -1, 1, 2},
			expected:  []int{-2, -1, 1, 2},
		},
		{
			asteroids: []int{-2, -2, 1, -2},
			expected:  []int{-2, -2, -2},
		},
		{
			asteroids: []int{1, -2, -2, -2},
			expected:  []int{-2, -2, -2},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, asteroidCollisionV1(test.asteroids), test.asteroids)
	}
}
