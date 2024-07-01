package main

func main() {
	// TODO: implement here
}

func isSubsequenceV1(s string, t string) bool {
	if len(s) > len(t) {
		return false
	} else if len(s) == 0 {
		return true
	}
	candidates := []byte(s)
	for _, char := range []byte(t) {
		if char == candidates[0] {
			candidates = candidates[1:]
		}
		if len(candidates) == 0 {
			return true
		}
	}
	return false
}
