package main

func main() {

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
