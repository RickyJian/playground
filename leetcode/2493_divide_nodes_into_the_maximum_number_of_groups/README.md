# 2493. Divide Nodes Into the Maximum Number of Groups

給定一個 n 個節點的無向圖，及一個 2 維陣列 edges([ai, bi]：表示 ai 和 bi 有一條雙向的邊)。

請將圖分成 m 組並滿足以下需求：

1. 圖中每個節點只屬於一組
2. 圖中每個邊連接的兩個點(ai, bi)，若 ai 屬於 x 組 bi 屬於 y 組，那麼 | x - y | = 1

請找出最大可將節點分為幾組，若無法找出請回傳 -1。

**注意圖可能是不連通的**

[2493_Divide_Nodes_Into_the_Maximum_Number_of_Groups](https://leetcode.com/problems/divide-nodes-into-the-maximum-number-of-groups/description/)
