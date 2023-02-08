package main

var tests = []*struct {
	n        int
	dislikes [][]int
	expected bool
}{
	{
		n:        4,
		dislikes: [][]int{{1, 2}, {1, 3}, {2, 4}},
		expected: true,
	},
	{
		n:        3,
		dislikes: [][]int{{1, 2}, {1, 3}, {2, 3}},
		expected: true,
	},
}
