package main

import (
	"fmt"
)

var (
	nums  = []int{1, 2, 3}
	nums2 = []int{4, 3, 2, 1}
	nums3 = []int{9}
)

func main() {
	fmt.Println(plusOne(nums))
	fmt.Println(plusOne(nums2))
	fmt.Println(plusOne(nums3))
}

func plusOne(digits []int) []int {
	digitLen := len(digits)
	results := make([]int, digitLen)
	var carry bool
	for i := digitLen - 1; ; i-- {
		if i < 0 && !carry {
			break
		}

		if i < 0 {
			results = append([]int{1}, results...)
			carry = false
		} else {
			digit := digits[i]
			if i == digitLen-1 || carry {
				if digit == 9 {
					digit = 0
					carry = true
				} else {
					digit++
					carry = false
				}
			}
			results[i] = digit
		}
	}
	return results
}
