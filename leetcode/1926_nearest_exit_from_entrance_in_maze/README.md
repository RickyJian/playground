# 1926. Nearest Exit from Entrance in Maze

## 題目描述

你被給予一個 `m x n` 的迷宮矩陣 `maze`（索引從 0 開始），其中：
- `maze[i][j] = .` 代表一個空位
- `maze[i][j] = +` 代表一面牆

你還被給予迷宮的入口位置 `entrance`，其中 `entrance = [row, col]` 表示你最初所在的格子位置。

一步移動是指從一個格子移動到相鄰的空位。你不能穿過牆壁，也不能走出迷宮的邊界。你的目標是找到從入口到最近出口的最少步數。

出口被定義為：
1. 位於迷宮邊界的空位
2. 不能是入口位置

如果無法到達任何出口，請回傳 `-1`。

## 範例

### 範例 1：
```
輸入：maze = [["+","+",".","+"],[".",".",".","+"],["+","+","+","."]], entrance = [1,2]
輸出：1
解釋：從入口 [1,2] 開始，我們可以一步到達出口 [1,0]。
```

### 範例 2：
```
輸入：maze = [["+","+","+"],[".",".","."],["+","+","+"]], entrance = [1,0]
輸出：2
解釋：從入口 [1,0] 開始，我們需要兩步才能到達出口 [1,2]。
```

### 範例 3：
```
輸入：maze = [[".","+"]], entrance = [0,0]
輸出：-1
解釋：從入口 [0,0] 開始，我們無法到達任何出口。
```

## 限制條件
- `maze.length == m`
- `maze[i].length == n`
- `1 <= m, n <= 100`
- `maze[i][j]` 是 `'.'` 或 `'+'`
- `entrance.length == 2`
- `0 <= entrancerow < m`
- `0 <= entrancecol < n`
- `entrance` 一定是空位

[1926. Nearest Exit from Entrance in Maze](https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/description/?envType=study-plan-v2&envId=leetcode-75)