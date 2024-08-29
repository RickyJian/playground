package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPredictPartyVictoryV1(t *testing.T) {
	var tests = []*struct {
		senate   string
		expected string
	}{
		{
			senate:   "RD",
			expected: "Radiant",
		},
		{
			senate:   "RDD",
			expected: "Dire",
		},
		{
			senate:   "RRR",
			expected: "Radiant",
		},
		{
			senate:   "DDRRR",
			expected: "Dire",
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, predictPartyVictoryV1(test.senate), test.senate)
	}
}
