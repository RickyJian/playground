package digit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigitCount1(t *testing.T) {
	tests := []struct {
		no       int
		expected int
	}{
		{
			no:       2,
			expected: 1,
		},
		{
			no:       12,
			expected: 2,
		},
		{
			no:       112,
			expected: 3,
		},
		{
			no:       1122,
			expected: 4,
		},
		{
			no:       1,
			expected: 1,
		},
		{
			no:       10,
			expected: 2,
		},
		{
			no:       100,
			expected: 3,
		},
		{
			no:       1000,
			expected: 4,
		},
	}
	for _, test := range tests {
		assert.Equal(t, DigitCount1(test.no), test.expected)
	}
}

func BenchmarkDigitCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DigitCount1(i)
	}
}

func TestDigitCount2(t *testing.T) {
	tests := []struct {
		no       int
		expected int
	}{
		{
			no:       2,
			expected: 1,
		},
		{
			no:       12,
			expected: 2,
		},
		{
			no:       112,
			expected: 3,
		},
		{
			no:       1122,
			expected: 4,
		},
		{
			no:       1,
			expected: 1,
		},
		{
			no:       10,
			expected: 2,
		},
		{
			no:       100,
			expected: 3,
		},
		{
			no:       1000,
			expected: 4,
		},
	}
	for _, test := range tests {
		assert.Equal(t, DigitCount2(test.no), test.expected)
	}
}

func BenchmarkDigitCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DigitCount2(i)
	}
}

func TestGetDigitValue1(t *testing.T) {
	tests := []struct {
		no       int
		digits   int
		expected []int
	}{
		{
			no:       1,
			digits:   1,
			expected: []int{1},
		},
		{
			no:       12,
			digits:   2,
			expected: []int{2, 1},
		},
		{
			no:       123,
			digits:   3,
			expected: []int{3, 2, 1},
		},
		{
			no:       1234,
			digits:   4,
			expected: []int{4, 3, 2, 1},
		},
	}
	for _, test := range tests {
		assert.Equal(t, GetDigitValue1(test.no, test.digits), test.expected)
	}
}

func BenchmarkGetDigitValue1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetDigitValue1(i, DigitCount1(i))
	}
}

func TestGetDigitValue2(t *testing.T) {
	tests := []struct {
		no       int
		digits   int
		expected []int
	}{
		{
			no:       1,
			digits:   1,
			expected: []int{1},
		},
		{
			no:       12,
			digits:   2,
			expected: []int{2, 1},
		},
		{
			no:       123,
			digits:   3,
			expected: []int{3, 2, 1},
		},
		{
			no:       1234,
			digits:   4,
			expected: []int{4, 3, 2, 1},
		},
	}
	for _, test := range tests {
		assert.Equal(t, GetDigitValue1(test.no, test.digits), test.expected)
	}
}

func BenchmarkGetDigitValue2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetDigitValue2(i, DigitCount1(i))
	}
}
