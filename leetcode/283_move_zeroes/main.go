package main

func main() {
	// TODO: implement here
}

func moveZeroesV1(nums []int) {
	var nonZeroIdx int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != 0 {
			nonZeroIdx++
			continue
		}

		next := nums[i+1]
		if next != 0 {
			nums[nonZeroIdx] = next
			nums[i+1] = 0
			nonZeroIdx++
		}
	}
}
