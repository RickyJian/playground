package main

func main() {
	// TODO: implement here
}

func largestAltitudeV1(gain []int) int {
	var prefix, max int
	for _, g := range gain {
		prefix += g
		if prefix > max {
			max = prefix
		}
	}
	return max
}
