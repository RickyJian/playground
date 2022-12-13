package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(magnificentSets(6, [][]int{{1, 2}, {1, 4}, {1, 5}, {2, 6}, {2, 3}, {4, 6}}))
	// fmt.Println(magnificentSetsBipartite(6, [][]int{{1, 2}, {1, 4}, {1, 5}, {2, 6}, {2, 3}, {4, 6}}))
	// fmt.Println(magnificentSetsAns(6, [][]int{{1, 2}, {1, 4}, {1, 5}, {2, 6}, {2, 3}, {4, 6}}))
}

func magnificentSets(n int, edges [][]int) int {
	// 1) 製圖：找出每個節點可以訪問的各節點
	//			2維陣列第一個陣列是每個節點，第二個陣列是每個節點可走訪的節點
	graph := make([][]int, n)
	for _, edge := range edges {
		// 由於題目節點是由 1 開始，但用 0 開始會易於後續處理，因此在這皆 -1
		n1, n2 := edge[0]-1, edge[1]-1
		graph[n1] = append(graph[n1], n2)
		graph[n2] = append(graph[n2], n1)
	}

	// 2) DFS：由於節點有可能不連通因此須將圖分組，待後續找出最大分組時使用
	var group [][]int
	visited := make(map[int]struct{})
	for i := 0; i < n; i++ {
		if _, ok := visited[i]; ok {
			// 走訪過的節點不再走訪
			continue
		}

		visited[i] = struct{}{}
		group = append(group, dfs(graph, i, visited))
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
			// 由於 bfs 出來的階層皆不同，須將群組中的節點皆走訪一次，找出最大的階層
			depth = int(math.Max(float64(depth), float64(count)))
		}
		result += depth
	}
	return result
}

// 將圖分組
func dfs(graph [][]int, n int, visited map[int]struct{}) []int {
	results := []int{n}
	for _, node := range graph[n] {
		if _, ok := visited[node]; ok {
			// 走訪過節點不須再走訪
			continue
		}

		visited[node] = struct{}{}
		results = append(results, dfs(graph, node, visited)...)
	}
	return results
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
					// 不符合題目要求，所有相依的節點都必須相鄰
					return -1
				} else if ok {
					// 走訪過節點不須再走訪
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
