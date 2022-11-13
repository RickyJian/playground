package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []*struct {
	robot    []int
	factory  [][]int
	expected int64
}{
	// {
	// 	robot: []int{0, 4, 6},
	// 	factory: [][]int{
	// 		{2, 2},
	// 		{6, 2},
	// 	},
	// 	expected: 4,
	// },
	// {
	// 	robot: []int{1, -1},
	// 	factory: [][]int{
	// 		{-2, 1},
	// 		{2, 1},
	// 	},
	// 	expected: 2,
	// },
	// {
	// 	robot: []int{9, 11, 99, 101},
	// 	factory: [][]int{
	// 		{10, 1},
	// 		{7, 1},
	// 		{14, 1},
	// 		{100, 1},
	// 		{96, 1},
	// 		{103, 1},
	// 	},
	// 	expected: 6,
	// },
	{
		robot: []int{79215383, 708490359, -779179404, 713376652, -368850098, 573013032, 195489859, 121470584, 916616893, 327266713, 950673412, 410723622, 538863648, 170740409, 753199490},
		factory: [][]int{
			{-284344805, 4},
			{349360740, 6},
			{-360820857, 2},
			{-493544411, 13},
			{-28182860, 4},
			{-117519725, 13},
			{-117519725, 13},
			{-294274103, 9},
		},
		expected: 5121930465,
	},
}

func TestMinimumTotalDistanceDFS(t *testing.T) {
	for _, test := range tests {
		assert.Equal(t, test.expected, minimumTotalDistanceDFS(test.robot, test.factory))
	}
}

// func TestMinimumTotalDistanceDFSMemo(t *testing.T) {
// 	for _, test := range tests {
// 		assert.Equal(t, test.expected, minimumTotalDistanceDFSMemo(test.robot, test.factory))
// 	}
// }
