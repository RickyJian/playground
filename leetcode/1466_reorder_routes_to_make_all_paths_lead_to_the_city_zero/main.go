package main

func main() {
	// TODO: implement here
}

func minReorder(n int, connections [][]int) int {
	// 作圖
	graph := make(map[int][][]int, len(connections)*2)
	for _, connection := range connections {
		graph[connection[0]] = append(graph[connection[0]], connection)
		graph[connection[1]] = append(graph[connection[1]], connection)
	}
	return minReorderDFS(graph, make([]bool, n), 0)
}

func minReorderDFS(graph map[int][][]int, visited []bool, target int) int {
	if visited[target] {
		return 0
	}

	visited[target] = true
	var count int
	for _, nodes := range graph[target] {
		var next int
		if nodes[0] == target {
			next = nodes[1]
			if !visited[next] {
				count++
			}
		} else {
			next = nodes[0]
		}
		count += minReorderDFS(graph, visited, next)
	}
	return count
}
