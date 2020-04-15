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
