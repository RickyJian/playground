package digit

import (
	"strconv"
)

// DigitCount1 count digits by loop
// best practice
func DigitCount1(no int) int {
	if no < 10 {
		return 1
	}

	digit := 1
	for i := 10; ; i = i * 10 {
		if no < i {
			break
		}
		digit++
	}
	return digit
}

// DigitCount2 count digits by strconv
func DigitCount2(no int) int {
	return len(strconv.Itoa(no))
}

// GetDigitValue1 get digit value by calculate, and array value from lsb to msb
func GetDigitValue1(no, digits int) []int {
	values := make([]int, digits)
	base := 10
	for i := 0; i < digits; i++ {
		values[i] = no % base / (base / 10)
		base *= 10
	}
	return values
}

// GetDigitValue2 get digit value by strconv, and array value from lsb to msb
// best practice
func GetDigitValue2(no, digits int) []int {
	noStr := strconv.Itoa(no)
	values := make([]int, digits)
	reverse := digits - 1
	for i := 0; i < digits; i++ {
		values[i], _ = strconv.Atoi(string(noStr[reverse]))
		reverse--
	}
	return values
}
