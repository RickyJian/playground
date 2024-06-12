package main

import (
	"strings"
)

func main() {
	// TODO: implement here
}

func reverseWords(s string) string {
	strs := strings.Split(s, " ")
	if len(strs) == 1 {
		return s
	}

	for i, j := 0, len(strs)-1; i < j; {
		prev, suff := strs[i], strs[j]
		if prev == " " || prev == "" {
			i++
		}
		if suff == " " || suff == "" {
			j--
		}
		if prev != "" && suff != "" {
			strs[i], strs[j] = strs[j], strs[i]
			i++
			j--
		}
	}
	var builder strings.Builder
	for _, str := range strs {
		if str == "" || str == " " {
			continue
		}
		builder.WriteString(str)
		builder.WriteString(" ")
	}
	return strings.TrimRight(builder.String(), " ")
}
