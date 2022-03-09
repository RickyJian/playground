package main

import (
	"testing"
)

func BenchmarkInitMapWithSize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		initMapWithSize(n)
	}
}

func BenchmarkInitMapWithoutSize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		initMapWithoutSize(n)
	}
}

func BenchmarkInitMapWithInterface(b *testing.B) {
	for n := 0; n < b.N; n++ {
		initMapWithInterface(n)
	}
}

func BenchmarkInitMapWithEmptyStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		initMapWithEmptyStruct(n)
	}
}
