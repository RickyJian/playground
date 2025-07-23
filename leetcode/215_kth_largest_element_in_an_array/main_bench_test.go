package main

import (
	"math/rand"
	"testing"
	"time"
)

// 生成測試數據
func generateTestData(size int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = rand.Intn(20000) - 10000 // 範圍 -10000 到 10000
	}
	return nums
}

// Benchmark findKthLargest (原始版本)
func BenchmarkFindKthLargest_Small(b *testing.B) {
	nums := generateTestData(100)
	k := 10

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findKthLargest(nums, k)
	}
}

func BenchmarkFindKthLargest_Medium(b *testing.B) {
	nums := generateTestData(1000)
	k := 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findKthLargest(nums, k)
	}
}

func BenchmarkFindKthLargest_Large(b *testing.B) {
	nums := generateTestData(5000)
	k := 500

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findKthLargest(nums, k)
	}
}

// Benchmark findKthLargestV2 (二分搜尋版本)
func BenchmarkFindKthLargestV2_Small(b *testing.B) {
	nums := generateTestData(100)
	k := 10

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findKthLargestV2(nums, k)
	}
}

func BenchmarkFindKthLargestV2_Medium(b *testing.B) {
	nums := generateTestData(1000)
	k := 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findKthLargestV2(nums, k)
	}
}

func BenchmarkFindKthLargestV2_Large(b *testing.B) {
	nums := generateTestData(5000)
	k := 500

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findKthLargestV2(nums, k)
	}
}

// 比較不同 k 值的效能
func BenchmarkFindKthLargest_K1(b *testing.B) {
	nums := generateTestData(1000)
	k := 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findKthLargest(nums, k)
	}
}

func BenchmarkFindKthLargestV2_K1(b *testing.B) {
	nums := generateTestData(1000)
	k := 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findKthLargestV2(nums, k)
	}
}

func BenchmarkFindKthLargest_KHalf(b *testing.B) {
	nums := generateTestData(1000)
	k := 500

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findKthLargest(nums, k)
	}
}

func BenchmarkFindKthLargestV2_KHalf(b *testing.B) {
	nums := generateTestData(1000)
	k := 500

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findKthLargestV2(nums, k)
	}
}

// 記憶體分配測試
func BenchmarkFindKthLargest_Memory(b *testing.B) {
	nums := generateTestData(1000)
	k := 100

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		findKthLargest(nums, k)
	}
}

func BenchmarkFindKthLargestV2_Memory(b *testing.B) {
	nums := generateTestData(1000)
	k := 100

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		findKthLargestV2(nums, k)
	}
}

// 最壞情況測試（已排序的陣列）
func BenchmarkFindKthLargest_WorstCase(b *testing.B) {
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = i // 升序排列
	}
	k := 500

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findKthLargest(nums, k)
	}
}

func BenchmarkFindKthLargestV2_WorstCase(b *testing.B) {
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = i // 升序排列
	}
	k := 500

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findKthLargestV2(nums, k)
	}
}

// 重複元素測試
func BenchmarkFindKthLargest_Duplicates(b *testing.B) {
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = i % 10 // 只有 0-9 的重複值
	}
	k := 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findKthLargest(nums, k)
	}
}

func BenchmarkFindKthLargestV2_Duplicates(b *testing.B) {
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = i % 10 // 只有 0-9 的重複值
	}
	k := 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findKthLargestV2(nums, k)
	}
}

// Benchmark findKthLargestV3 (新版本)
func BenchmarkFindKthLargestV3_Small(b *testing.B) {
	nums := generateTestData(100)
	k := 10

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findKthLargestV3(nums, k)
	}
}

func BenchmarkFindKthLargestV3_Medium(b *testing.B) {
	nums := generateTestData(1000)
	k := 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findKthLargestV3(nums, k)
	}
}

func BenchmarkFindKthLargestV3_Large(b *testing.B) {
	nums := generateTestData(5000)
	k := 500

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findKthLargestV3(nums, k)
	}
}

func BenchmarkFindKthLargestV3_K1(b *testing.B) {
	nums := generateTestData(1000)
	k := 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findKthLargestV3(nums, k)
	}
}

func BenchmarkFindKthLargestV3_KHalf(b *testing.B) {
	nums := generateTestData(1000)
	k := 500

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findKthLargestV3(nums, k)
	}
}

func BenchmarkFindKthLargestV3_Memory(b *testing.B) {
	nums := generateTestData(1000)
	k := 100

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		findKthLargestV3(nums, k)
	}
}

func BenchmarkFindKthLargestV3_WorstCase(b *testing.B) {
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = i // 升序排列
	}
	k := 500

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findKthLargestV3(nums, k)
	}
}

func BenchmarkFindKthLargestV3_Duplicates(b *testing.B) {
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = i % 10 // 只有 0-9 的重複值
	}
	k := 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findKthLargestV3(nums, k)
	}
}
