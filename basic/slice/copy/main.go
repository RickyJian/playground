package main

func main() {

}

func SliceCopy1(new, old []string) []string {
	c := make([]string, len(new)+len(old))
	for i, v := range new {
		c[i] = v
	}
	for i, v := range old {
		c[len(new)+i] = v
	}
	return c
}

func SliceCopy2(new, old []string) []string {
	c := make([]string, len(new)+len(old))
	copy(c, new)
	copy(c[len(new):], old)
	return c
}
