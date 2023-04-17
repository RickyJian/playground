package main

import "testing"

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
