package main

func main() {
}

func networkDelayTime(times [][]int, n int, k int) int {
	nodes := make(map[int][][]int)
	for _, time := range times {
		nodes[time[0]] = append(nodes[time[0]], []int{time[1], time[2]})
	}

	minPath := make(map[int]int)
	reached := make(map[int]struct{})
	dfs(nodes, reached, minPath, k, 0)
	if len(reached) != n {
		return -1
	}

	var max int
	for _, path := range minPath {
		if max < path {
			max = path
		}
	}
	return max
}

func dfs(nodes map[int][][]int, reached map[int]struct{}, minPath map[int]int, from, weight int) {
	reached[from] = struct{}{}
	if minWeight, ok := minPath[from]; ok && minWeight <= weight {
		return
	}

	minPath[from] = weight
	for _, children := range nodes[from] {
		dfs(nodes, reached, minPath, children[0], children[1]+weight)
	}
}
