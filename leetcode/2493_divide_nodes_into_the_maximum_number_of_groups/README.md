# 2493. Divide Nodes Into the Maximum Number of Groups

图中每个节点都只属于一个组。
图中每条边连接的两个点[ai, bi]，如果 ai 属于编号为 x 的组，bi 属于编号为 y 的组，那么 |y - x| = 1 。
请你返回最多可以将节点分为多少个组（也就是最大的 m ）。如果没办法在给定条件下分组，请你返回 -1 。

給定一個 n 個節點的無向圖，及一個 2 維陣列 edges([ai, bi]：表示 ai 和 bi 有一條雙向的邊)。

請將圖分成 m 組並滿足以下需求：

1. 圖中每個節點只屬於一組
2. 途中每個邊連接的兩個點(ai, bi)，若 ai 屬於 x 組 bi 屬於 y 組，那麼 | x - y | = 1
3. 請找出最大可將節點分為幾組，若無法找出請回傳 -1。

**注意圖可能是不連通的**

[2493_Divide_Nodes_Into_the_Maximum_Number_of_Groups](https://leetcode.com/problems/divide-nodes-into-the-maximum-number-of-groups/description/)
