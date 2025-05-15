package main

func main() {
	// TODO: implement here
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	if len(edges) <= 1 {
		return true
	}

	// 作圖（找出每個 node 相鄰的 node）
	grapth := make(map[int][]int, n)
	for _, edge := range edges {
		grapth[edge[0]] = append(grapth[edge[0]], edge[1])
		grapth[edge[1]] = append(grapth[edge[1]], edge[0])
	}
	// 走訪所有的 node
	return validPathDFS(grapth, make([]bool, n), source, destination)
}

func validPathDFS(grapth map[int][]int, visited []bool, from, destination int) bool {
	if from == destination {
		return true
	} else if visited[from] {
		// 已訪問節點不須再次造訪
		return false
	}

	visited[from] = true
	for _, next := range grapth[from] {
		if validPathDFS(grapth, visited, next, destination) {
			return true
		}
	}
	return false
}
