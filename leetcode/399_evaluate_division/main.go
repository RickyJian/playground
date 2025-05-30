package main

func main() {
	// TODO: implement here
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string][]*node)
	for i, equation := range equations {
		word1, word2 := equation[0], equation[1]
		graph[word1] = append(graph[word1], &node{next: word2, weight: values[i]})
		graph[word2] = append(graph[word2], &node{next: word1, weight: 1 / values[i]})
	}
	results := make([]float64, len(queries))
	for i, q := range queries {
		word1, word2 := q[0], q[1]
		if len(graph[word1]) == 0 || len(graph[word2]) == 0 {
			results[i] = -1.0
		} else if word1 == word2 {
			results[i] = 1.0
		} else {
			results[i] = dfs(graph, make(map[string]bool), word1, word2)
		}
	}
	return results
}

func dfs(graph map[string][]*node, visited map[string]bool, from, to string) float64 {
	if from == to {
		return 1.0
	}

	visited[from] = true
	for _, node := range graph[from] {
		if visited[node.next] {
			continue
		} else if result := dfs(graph, visited, node.next, to); result != -1.0 {
			return result * node.weight
		}
	}
	return -1.0
}

type node struct {
	next   string
	weight float64
}
