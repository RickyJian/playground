package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkNewSliceWithMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewSliceWithMake(i)
	}
}

func BenchmarkNewSliceWithMakeAndSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewSliceWithMakeAndSize(i)
	}
}

func BenchmarkNewSliceWithoutMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewSliceWithoutMake(i)
	}
}

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

func TestPrepend1(t *testing.T) {
	tests := []struct {
		no       int
		nos      []int
		expected []int
	}{
		{
			no:       1,
			nos:      []int{2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, Prepend1(test.no, test.nos))
	}
}

func TestPrepend2(t *testing.T) {
	tests := []struct {
		no       int
		nos      []int
		expected []int
	}{
		{
			no:       1,
			nos:      []int{2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, Prepend2(test.no, test.nos))
	}
}

func BenchmarkPrepend1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Prepend1(i, make([]int, i+1))
	}
}

func BenchmarkPrepend2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Prepend2(i, make([]int, i+1))
	}
}
