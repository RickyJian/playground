package main

func main() {
	// TODO: implement here
}

func findDifferenceV1(nums1 []int, nums2 []int) [][]int {
	set1 := map[int]struct{}{}
	for _, num := range nums1 {
		set1[num] = struct{}{}
	}
	results2 := make([]int, 0)
	set2 := map[int]struct{}{}
	for _, num := range nums2 {
		if _, ok := set1[num]; !ok {
			results2 = append(results2, num)
			set1[num] = struct{}{}
		}
		set2[num] = struct{}{}
	}
	results1 := make([]int, 0)
	for _, num := range nums1 {
		if _, ok := set2[num]; !ok {
			results1 = append(results1, num)
			set2[num] = struct{}{}
		}
	}
	return [][]int{results1, results2}
}

func findDifferenceV2(nums1 []int, nums2 []int) [][]int {
	const count = 2001
	h1 := make([]bool, count)
	for _, num := range nums1 {
		h1[num+1000] = true
	}
	results2 := make([]int, 0)
	h2 := make([]bool, count)
	for _, num := range nums2 {
		nn := num + 1000
		if !h1[nn] {
			results2 = append(results2, num)
			h1[nn] = true
		}
		h2[nn] = true
	}
	results1 := make([]int, 0)
	for _, num := range nums1 {
		if nn := num + 1000; !h2[nn] {
			results1 = append(results1, num)
			h2[nn] = true
		}
	}
	return [][]int{results1, results2}
}
