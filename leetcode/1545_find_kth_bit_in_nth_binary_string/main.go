package main

func main() {
}

func findKthBit(n int, k int) byte {
	if n == 1 {
		return '0'
	}

	length := (1 << n) - 1
	mid := length/2 + 1
	if mid == k {
		// mid == k：方程式中 `1` 值，直接回傳 `1`
		return '1'
	} else if k > mid {
		// k > mid：代表後續位置皆為 k < mid 的鏡像，須反轉回傳為來的值
		if findKthBit(n-1, length-k+1) == '0' {
			return '1'
		}
		return '0'
	}
	// k < mid：原始位置不須反轉，直接還傳找到的值
	return findKthBit(n-1, k)
}
