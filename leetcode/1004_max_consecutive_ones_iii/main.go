package main

func main() {
	// TODO: implement here
}

func longestOnesV1(nums []int, k int) int {
	if k >= len(nums) {
		return len(nums)
	}

	start := -1
	flipIdxes := make([]int, 0, k)
	var results int
	for i, num := range nums {
		if num == 0 {
			// check can flip, if not reset index
			if k == 0 && len(flipIdxes) == 0 {
				start = -1
			} else {
				// can flip
				flipIdxes = append(flipIdxes, i)
				if k > 0 {
					k--
				} else {
					start = flipIdxes[0] + 1
					flipIdxes = flipIdxes[1:]
				}
			}
		} else {
			if start == -1 {
				// init position
				start = i
			}
		}
		if start > -1 {
			if size := i - start + 1; results < size {
				results = size
			}
		}
	}
	return results
}
