package main

import (
	"strconv"
)

func main() {
	// TODO: implement here
}

func compressV1(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}
	begin, end := -1, -1
	var nc []byte
	var count int
	for i := 0; i < len(chars); i++ {
		current := chars[i]
		next := chars[i]
		if i < len(chars)-1 {
			next = chars[i+1]
		}
		if begin == -1 {
			begin = i
		}
		if current != next || i == len(chars)-1 {
			end = i + 1

			count++
			nc = append(nc, chars[begin])

			if end-begin > 1 {
				quantity := strconv.Itoa(len(chars[begin:end]))
				count += len(quantity)
				for _, q := range quantity {
					nc = append(nc, byte(q))
				}
			}
			begin, end = -1, -1
		}
	}
	chars = chars[:len(nc)]
	copy(chars, nc)
	return count
}
