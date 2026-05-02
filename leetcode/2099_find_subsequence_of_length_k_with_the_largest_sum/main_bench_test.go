package main

import (
	"math/rand"
	"testing"
	"time"
)

func generateBenchData(size int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = rand.Intn(200000) - 100000 // 範圍 -100000 到 100000
	}
	return nums
}

func BenchmarkMaxSubsequence_Small(b *testing.B) {
	nums := generateBenchData(100)
	k := 10

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxSubsequence(nums, k)
	}
}

func BenchmarkMaxSubsequence_Medium(b *testing.B) {
	nums := generateBenchData(500)
	k := 50

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxSubsequence(nums, k)
	}
}

func BenchmarkMaxSubsequence_Large(b *testing.B) {
	nums := generateBenchData(1000)
	k := 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxSubsequence(nums, k)
	}
}
