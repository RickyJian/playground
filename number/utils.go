package number

// DigitCount1 count digits by loop
func DigitCount1(no int) int {
	if no < 10 {
		return 1
	}

	digit := 1
	for i := 10; ; i = i * 10 {
		if no <= i {
			break
		}
		digit++
	}
	return digit
}
