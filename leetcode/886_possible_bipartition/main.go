package main

func main() {
	// TODO: implement here
}

// DFS solution
// refer to: https://en.wikipedia.org/wiki/Bipartite_graph
func possibleBipartitionDFS(n int, dislikes [][]int) bool {
	graph := make([][]int, n)
	for _, nodes := range dislikes {
		n1, n2 := nodes[0]-1, nodes[1]-1
		graph[n1] = append(graph[n1], n2)
		graph[n2] = append(graph[n2], n1)
	}

	colors := make([]int, n)
	for i := 0; i < n; i++ {
		if colors[i] == 0 {
			if !dfs(graph, i, 1, colors) {
				return false
			}
		}
	}
	return true
}

func dfs(graph [][]int, node int, color int, colors []int) bool {
	colors[node] = color
	for _, nextNode := range graph[node] {
		if nextColor := colors[nextNode]; nextColor == 0 {
			if !dfs(graph, nextNode, -color, colors) {
				return false
			}
		} else if nextColor == color {
			return false
		}
	}
	return true
}

// Union Find solution
func possibleBipartitionUF(n int, dislikes [][]int) bool {
	return false
}
