package main

func main() {
	// TODO: implement here
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var max int
	for _, candy := range candies {
		if max < candy {
			max = candy
		}
	}
	results := make([]bool, len(candies))
	for i, candy := range candies {
		if max <= candy+extraCandies {
			results[i] = true
		}
	}
	return results
}
