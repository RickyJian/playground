package main

func main() {
	// TODO: implement here
}

// DFS solution
// refer to: https://en.wikipedia.org/wiki/Bipartite_graph
func possibleBipartitionDFS(n int, dislikes [][]int) bool {
	// 1) 製圖
	graph := make([][]int, n)
	for _, nodes := range dislikes {
		n1, n2 := nodes[0]-1, nodes[1]-1
		graph[n1] = append(graph[n1], n2)
		graph[n2] = append(graph[n2], n1)
	}

	// 2) 染色
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
	graph := make([][]int, n)
	for _, dislike := range dislikes {
		p1, p2 := dislike[0]-1, dislike[1]-1
		graph[p1] = append(graph[p1], p2)
		graph[p2] = append(graph[p2], p1)
	}

	u := newUF(n)
	for i := 0; i < n; i++ {
		for _, j := range graph[i] {
			// 若 i, j 查找到相同的 root 代表無法將齊分為兩群
			if u.find(i) == u.find(j) {
				return false
			}
			// 以第一個人作為代表合併其餘人
			u.union(graph[i][0], j)
		}
	}
	return true
}

type uf []int

func newUF(n int) *uf {
	u := make(uf, n)
	for i := 0; i < n; i++ {
		u[i] = i
	}
	return &u
}

func (u uf) union(i, j int) {
	findI, findJ := u.find(i), u.find(j)
	if findI != findJ {
		u[findJ] = findI
	}
}

func (u uf) find(idx int) int {
	for u[idx] != idx {
		idx = u[idx]
	}
	return idx
}
