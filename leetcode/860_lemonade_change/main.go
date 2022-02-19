package main

import (
	"fmt"
)

func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 20}))
}

func lemonadeChange(bills []int) bool {
	var five, ten int
	for _, bill := range bills {
		switch bill {
		case 5:
			five++
		case 10:
			if five == 0 {
				return false
			}
			ten++
			five--
		case 20:
			if five == 0 || (ten == 0 && five < 3) {
				return false
			}

			if ten > 0 {
				ten--
				five--
			} else if ten == 0 {
				five -= 3
			}
		}
	}
	return true
}
