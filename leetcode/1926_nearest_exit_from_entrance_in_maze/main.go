package main

func main() {
	// 測試程式碼將在 main_test.go 中實作
}

// nearestExit 尋找從入口到最近出口的最少步數
// maze: 迷宮矩陣，其中 '.' 代表空位，'+' 代表牆
// entrance: 入口位置 [row, col]
// 回傳：到達最近出口的步數，如果無法到達則回傳 -1
func nearestExit(maze [][]byte, entrance []int) int {
	// 走一次迷宮並完成 graph
	const emptyHole = '.'
	nodes := make(map[node][]node)
	for row, cols := range maze {
		for col, cell := range cols {
			if emptyHole == cell {
				fromNode := node{row: row, col: col}
				toNodes := make([]node, 0, 4)
				if row > 0 && maze[row-1][col] == emptyHole {
					// 向上
					toNodes = append(toNodes, node{row: row - 1, col: col})
				}
				if row < len(maze)-1 && maze[row+1][col] == emptyHole {
					// 向下
					toNodes = append(toNodes, node{row: row + 1, col: col})
				}
				if col > 0 && maze[row][col-1] == emptyHole {
					// 向左
					toNodes = append(toNodes, node{row: row, col: col - 1})
				}
				if col < len(cols)-1 && maze[row][col+1] == emptyHole {
					// 向右
					toNodes = append(toNodes, node{row: row, col: col + 1})
				}
				nodes[fromNode] = append(nodes[fromNode], toNodes...)
			}
		}
	}

	// 無路可走
	start := node{row: entrance[0], col: entrance[1]}
	if len(nodes[start]) == 0 {
		return -1
	}

	// BFS 找出最短路徑
	visited := make(map[node]bool)
	levels := make(map[node]int)
	levels[start] = 0
	queue := []node{start}
	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		if visited[first] {
			continue
		}

		visited[first] = true
		for _, nextNode := range nodes[first] {
			if visited[nextNode] {
				continue
			}

			levels[nextNode] = levels[first] + 1
			if nextNode.row == 0 || nextNode.col == 0 || nextNode.row == len(maze)-1 || nextNode.col == len(maze[0])-1 {
				return levels[nextNode]
			}
			queue = append(queue, nextNode)
		}
	}
	return -1
}

type node struct {
	row int
	col int
}

// 優化方向：
//  1. 減少製圖：直接 BFS 當走訪到出口一定是最短路徑，因此不需要全部 node 都走過一遍
//  2. 使用 `[][3]int` 代替 `[]node`
//  3. visited使用 `[][]bool` 代替 `map[node]bool`：由於是已知數量的矩陣（row x col），使用 map 會有不需要的雜湊計算和碰撞發生
func nearestExitOptimized(maze [][]byte, entrance []int) int {
	rows := len(maze)
	if rows == 0 {
		return -1
	}

	cols := len(maze[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	// [3]int:
	//  1. index 0: row
	//  2. index 1: col
	//  3. index 2: steps
	queue := [][3]int{{entrance[0], entrance[1], 0}}
	visited[entrance[0]][entrance[1]] = true
	// BFS
	dirs := [][2]int{
		{-1, 0}, // 上
		{1, 0},  // 下
		{0, -1}, // 左
		{0, 1},  // 右
	}
	const emptyHole = '.'
	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]

		// 檢查是否為出口（在邊界且不可為入口）
		row, col, steps := first[0], first[1], first[2]
		if (row == 0 || col == 0 || row == rows-1 || col == cols-1) &&
			(row != entrance[0] || col != entrance[1]) {
			return steps
		}

		// 訪問四周節點
		for _, dir := range dirs {
			newRow, newCol := row+dir[0], col+dir[1]
			if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
				// 邊界檢查
				continue
			} else if visited[newRow][newCol] {
				// 已走訪不再走訪
				continue
			} else if maze[newRow][newCol] != emptyHole {
				// 牆壁不可走訪
				continue
			}
			visited[newRow][newCol] = true
			queue = append(queue, [3]int{newRow, newCol, steps + 1})
		}
	}
	return -1
}

// 優化方向：
//  1. 移除 visited：走訪過得 hole 直接更改成 wall
func nearestExitOptimized2(maze [][]byte, entrance []int) int {
	rows := len(maze)
	if rows == 0 {
		return -1
	}
	cols := len(maze[0])
	dirs := [][2]int{
		{-1, 0}, // 上
		{1, 0},  // 下
		{0, -1}, // 左
		{0, 1},  // 右
	}

	const wall = '+'
	var steps int
	queue := [][2]int{{entrance[0], entrance[1]}}
	maze[entrance[0]][entrance[1]] = wall
	for len(queue) > 0 {
		for _, first := range queue {
			queue = queue[1:]
			row, col := first[0], first[1]
			if (row == 0 || col == 0 || row == rows-1 || col == cols-1) &&
				(row != entrance[0] || col != entrance[1]) {
				return steps
			}

			for _, dir := range dirs {
				newRow, newCol := row+dir[0], col+dir[1]
				if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
					// 邊界檢查
					continue
				} else if maze[newRow][newCol] == wall {
					// 牆壁不可走訪
					continue
				}
				// 走訪過的 node 標記成 `+`
				maze[newRow][newCol] = wall
				queue = append(queue, [2]int{newRow, newCol})
			}
		}
		steps++
	}
	return -1
}
