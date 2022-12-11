package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(magnificentSets(6, [][]int{{1, 2}, {1, 4}, {1, 5}, {2, 6}, {2, 3}, {4, 6}}))
	fmt.Println(magnificentSetsBipartite(6, [][]int{{1, 2}, {1, 4}, {1, 5}, {2, 6}, {2, 3}, {4, 6}}))
	fmt.Println(magnificentSetsAns(6, [][]int{{1, 2}, {1, 4}, {1, 5}, {2, 6}, {2, 3}, {4, 6}}))
}

// TLE
func magnificentSets(n int, edges [][]int) int {
	// 1) 製圖：找出每個節點可以訪問的各節點
	graph := make([][]int, n)
	for _, edge := range edges {
		n1, n2 := edge[0]-1, edge[1]-1
		graph[n1] = append(graph[n1], n2)
		graph[n2] = append(graph[n2], n1)
	}

	// 2) DFS：由於節點有可能不連通，因此須將圖分組
	var group [][]int
	visited := make(map[int]struct{})
	for i := 0; i < n; i++ {
		var nodes []int
		if _, ok := visited[i]; !ok {
			dfs(graph, i, &nodes, visited)
		}
		group = append(group, nodes)
	}

	// 3) BFS：檢查圖是否合法，並找最大階層
	var result int
	for _, nodes := range group {
		var depth int
		for _, node := range nodes {
			count := bfs(graph, node)
			if count == -1 {
				return -1
			}
			depth = int(math.Max(float64(depth), float64(count)))
		}
		result += depth
	}
	return result
}

// 將圖分組
func dfs(graph [][]int, n int, nodes *[]int, visited map[int]struct{}) {
	*nodes = append(*nodes, n)
	for _, node := range graph[n] {
		if _, ok := visited[node]; ok {
			continue
		}

		visited[node] = struct{}{}
		*nodes = append(*nodes, node)
		dfs(graph, node, nodes, visited)
	}
	return
}

// 檢查圖是否合法，並找最大階層
func bfs(graph [][]int, node int) int {
	var result int
	visited := make(map[int]int)
	queue := []int{node}
	for len(queue) > 0 {
		for _, n := range queue {
			queue = queue[1:]
			visited[n] = result
			for _, nextNode := range graph[n] {
				if level, ok := visited[nextNode]; ok && math.Abs(float64(result-level)) != 1 {
					return -1
				} else if _, ok := visited[nextNode]; ok {
					continue
				}

				queue = append(queue, nextNode)
				visited[nextNode] = result + 1
			}
		}
		result++
	}
	return result
}

func magnificentSetsBipartite(n int, edges [][]int) (ans int) {
	// 1) 製圖：找出每個節點可以訪問的各節點
	graph := make([][]int, n)
	for _, edge := range edges {
		n1, n2 := edge[0]-1, edge[1]-1
		graph[n1] = append(graph[n1], n2)
		graph[n2] = append(graph[n2], n1)
	}

	var result int
	colors := make([]int8, n)
	for i, c := range colors {
		if c == 0 {
			nodes, ok := dfsBipartite(graph, colors, i, 1)
			if !ok {
				return -1
			}
			var depth int
			for _, node := range nodes {
				depth = int(math.Max(float64(depth), float64(bfsBipartite(graph, node))))
			}
			result += depth
		}
	}
	return result
}

func dfsBipartite(graph [][]int, colors []int8, node int, c int8) ([]int, bool) {
	colors[node] = c
	nodes := []int{node}
	for _, y := range graph[node] {
		if colors[y] == c {
			return []int{}, false
		}

		if colors[y] == 0 {
			visitNodes, ok := dfsBipartite(graph, colors, y, -c)
			if !ok {
				return []int{}, false
			}
			nodes = append(nodes, visitNodes...)
		}
	}
	return nodes, true
}

func bfsBipartite(graph [][]int, node int) int {
	var result int
	visited := make(map[int]struct{})
	queue := []int{node}
	for len(queue) > 0 {
		for _, n := range queue {
			queue = queue[1:]
			for _, y := range graph[n] {
				if _, ok := visited[y]; !ok { // 没有在同一次 BFS 中访问过
					queue = append(queue, y)
				}
				visited[y] = struct{}{}
			}
			visited[n] = struct{}{}
		}
		result++
	}
	return result
}

// best solution
// refer to: https://leetcode.cn/problems/divide-nodes-into-the-maximum-number-of-groups/solution/mei-ju-qi-dian-pao-bfs-by-endlesscheng-s5bu/
func magnificentSetsAns(n int, edges [][]int) (ans int) {
	// 製圖：將每個 node 可到達的頂點
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0]-1, e[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	colors := make([]int8, n)
	var nodes []int
	var isBipartite func(int, int8) bool
	isBipartite = func(x int, c int8) bool { // 二分图判定，原理见视频讲解
		nodes = append(nodes, x)
		colors[x] = c
		for _, y := range g[x] {
			if colors[y] == c || colors[y] == 0 && !isBipartite(y, -c) {
				return false
			}
		}
		return true
	}

	time := make([]int, n) // 充当 vis 数组的作用（避免在 BFS 内部重复创建 vis 数组）
	clock := 0
	bfs := func(start int) (depth int) { // 返回从 start 出发的最大深度
		clock++
		time[start] = clock
		for q := []int{start}; len(q) > 0; depth++ {
			tmp := q
			q = nil
			for _, x := range tmp {
				for _, y := range g[x] {
					if time[y] != clock { // 没有在同一次 BFS 中访问过
						time[y] = clock
						q = append(q, y)
					}
				}
			}
		}
		return
	}

	for i, c := range colors {
		if c == 0 {
			nodes = nil
			if !isBipartite(i, 1) {
				return -1
			}
			maxDepth := 0
			for _, x := range nodes {
				maxDepth = int(math.Max(float64(maxDepth), float64(bfs(x))))
			}
			ans += maxDepth
		}
	}
	return
}
