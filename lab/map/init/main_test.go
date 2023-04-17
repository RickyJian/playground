package main

import (
	"testing"
)

func BenchmarkInitMapWithSize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		initMapWithSize(300000)
	}
}

func BenchmarkInitMapWithoutSize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		initMapWithoutSize(300000)
	}
}

func BenchmarkInitMapWithInterface(b *testing.B) {
	for n := 0; n < b.N; n++ {
		initMapWithInterface(300000)
	}
}

func BenchmarkInitMapWithEmptyStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		initMapWithEmptyStruct(300000)
	}
}
