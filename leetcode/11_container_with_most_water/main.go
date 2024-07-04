package main

func main() {
	// TODO: implement here
}

// Time Limit Exceeded
func maxAreaV1(height []int) int {
	var max int
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			highest := height[i]
			if height[j] < highest {
				highest = height[j]
			}
			if current := (j - i) * highest; current > max {
				max = current
			}
		}
	}
	return max
}

func maxAreaV2(height []int) int {
	var max int
	for s, e := 0, len(height)-1; s < e; {
		left, right := height[s], height[e]
		lowest := left
		if right < lowest {
			lowest = right
		}
		if current := (e - s) * lowest; current > max {
			max = current
		}
		if left <= right {
			s++
		} else {
			e--
		}

	}
	return max
}
