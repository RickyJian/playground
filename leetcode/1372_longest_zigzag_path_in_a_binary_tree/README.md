# 1372. Longest ZigZag Path in a Binary Tree

給定一個二元樹，請以 `ZigZag` 走法走訪二元樹的節點，

並根據以下規則找出最長的路徑，

> ZigZag 走訪長度：訪問的節點數量減 1

規則：

1. 選定任一節點和方向（左或右）
2. 如果當前方向是右，則移動到當前節點的右子節點。否則移動到左子節點。
3. 將方向從右改為左，或從左改為右
4. 重複第二步和第三步，直到無法在樹中移動為止。

---

[1372. Longest ZigZag Path in a Binary Tree](https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree/)