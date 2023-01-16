# 547. Number of Provinces

有 n 個城市某些是相連某些沒有，若 a 城市相連 b 城市且 b 城市相連 c 城市，代表 a 間接連在 c 城市。

給定一個 n x n 的矩陣 isConnected：

1. isConnected[i][j] = 1：代表第 i 城市與第 j 城市相接
2. isConnected[i][j] = 0：代表第 i 城市與第 j 城市不相接

省份代表相連或間接相連的城市，請找出總共有幾個省份

[547_Number_of_Provinces](https://leetcode.com/problems/number-of-provinces/description/?q=union&orderBy=most_relevant&languageTags=golang)
