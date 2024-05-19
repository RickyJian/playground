package main

func main() {
	// TODO: implement here
}

func gcdOfStrings(str1 string, str2 string) string {
	size := len(str1)
	if str2Len := len(str2); str2Len < size {
		size = str2Len
	}
	matchLen := -1
	for i := 0; i <= size; i++ {
		if str1[:i] != str2[:i] {
			break
		}
		matchLen++
	}
LOOP:
	for matchLen > 0 {
		if len(str1)%matchLen > 0 || len(str2)%matchLen > 0 {
			matchLen--
			continue
		}
		prefix := str1[:matchLen]
		for i := matchLen; i < len(str1); i += matchLen {
			if str1[i:i+matchLen] != prefix {
				matchLen--
				continue LOOP
			}
		}
		for i := matchLen; i < len(str2); i += matchLen {
			if str2[i:i+matchLen] != prefix {
				matchLen--
				continue LOOP
			}
		}
		return prefix
	}
	return ""
}
