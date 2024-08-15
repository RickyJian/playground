package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveStars(t *testing.T) {
	var tests = []*struct {
		s        string
		expected string
	}{
		// {
		// 	s:        "leet**cod*e",
		// 	expected: "lecoe",
		// },
		{
			s:        "erase*****",
			expected: "",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, removeStarsV1(test.s), test.s)
	}
}
