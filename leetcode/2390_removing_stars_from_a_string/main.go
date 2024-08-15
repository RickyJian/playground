package main

func main() {
	// TODO: implement here
}

func removeStarsV1(s string) string {
	stack := make([]rune, 0, len(s))
	for _, word := range s {
		if word == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, word)
		}
	}
	return string(stack)
}
