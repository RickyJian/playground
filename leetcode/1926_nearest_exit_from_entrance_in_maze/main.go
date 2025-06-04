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
