package slice

// NewSliceWithMake add element use make
// best practice
func NewSliceWithMake(no int) {
	nos := make([]int, no)
	for i := 0; i < no; i++ {
		nos[i] = i
	}
}

// NewSliceWithMakeAndSize add element make memory size
func NewSliceWithMakeAndSize(no int) {
	nos := make([]int, 0, no)
	for i := 0; i < no; i++ {
		nos = append(nos, i)
	}

}

// NewSliceWithoutMake add element
func NewSliceWithoutMake(no int) {
	var nos []int
	for i := 0; i < no; i++ {
		nos = append(nos, i)
	}
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

func Prepend1(no int, nos []int) []int {
	return append([]int{no}, nos...)
}

func Prepend2(no int, nos []int) []int {
	nos = append(nos, 0)
	copy(nos[1:], nos)
	nos[0] = no
	return nos
}
