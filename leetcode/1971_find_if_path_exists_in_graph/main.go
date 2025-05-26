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

func validPathLoopDFS(n int, edges [][]int, source int, destination int) bool {
	grapth := make(map[int][]int, n)
	for _, edge := range edges {
		grapth[edge[0]] = append(grapth[edge[0]], edge[1])
		grapth[edge[1]] = append(grapth[edge[1]], edge[0])
	}

	visited := make([]bool, n)
	visited[source] = true
	stack := []int{source}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current == destination {
			return true
		}
		for _, next := range grapth[current] {
			if visited[next] {
				continue
			}
			visited[next] = true
			stack = append(stack, next)
		}
	}
	return false
}

func validPathBFS(n int, edges [][]int, source int, destination int) bool {
	grapth := make(map[int][]int, n)
	for _, edge := range edges {
		grapth[edge[0]] = append(grapth[edge[0]], edge[1])
		grapth[edge[1]] = append(grapth[edge[1]], edge[0])
	}

	queue := []int{source}
	visited := make([]bool, n)
	visited[source] = true
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current == destination {
			return true
		}
		for _, next := range grapth[current] {
			if visited[next] {
				continue
			}
			visited[next] = true
			queue = append(queue, next)
		}
	}
	return false
}

// best solution
func validPathUnionFind(n int, edges [][]int, source int, destination int) bool {
	uf := &unionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
	}
	// init node
	for i := range uf.parent {
		uf.parent[i] = i
		uf.rank[i] = 1 // 初始化 rank
	}
	for _, nodes := range edges {
		uf.union(nodes[0], nodes[1])
	}
	return uf.find(source) == uf.find(destination)
}

type unionFind struct {
	parent []int
	rank   []int
}

func (u *unionFind) union(from, to int) {
	findF, findT := u.find(from), u.find(to)
	if findF != findT {
		if u.rank[findF] < u.rank[findT] {
			u.parent[findF] = findT
		} else if u.rank[findF] > u.rank[findT] {
			u.parent[findT] = findF
		} else {
			u.parent[findT] = findF
			u.rank[findF]++
		}
	}
}

func (u *unionFind) find(target int) int {
	if u.parent[target] != target {
		u.parent[target] = u.find(u.parent[target])
	}
	return u.parent[target]
}
