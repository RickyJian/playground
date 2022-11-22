package main

func main() {
}

func networkDelayTimeDFS(times [][]int, n int, k int) int {
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

func networkDelayTimeBFS(times [][]int, n int, k int) int {
	children := map[int][][]int{}
	for _, time := range times {
		children[time[0]] = append(children[time[0]], []int{time[1], time[2]})
	}

	seen := map[int]bool{}
	minPath := map[int]int{}
	nodes := [][]int{{k, 0}}
	for len(nodes) != 0 {
		parentLen := len(nodes)
		for _, node := range nodes {
			from, weight := node[0], node[1]
			seen[from] = true
			if currentWeight, ok := minPath[from]; ok && weight >= currentWeight {
				continue
			}

			minPath[from] = weight
			for _, childNode := range children[from] {
				nodes = append(nodes, []int{childNode[0], weight + childNode[1]})
			}
		}
		copy(nodes, nodes[parentLen:])
		nodes = nodes[:len(nodes[parentLen:])]
	}

	if len(seen) != n {
		return -1
	}

	max := -1
	for _, path := range minPath {
		if max < path {
			max = path
		}
	}
	return max
}

// TODO: add Dijkstra Algorithm solution
