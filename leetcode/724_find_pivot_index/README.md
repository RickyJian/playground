# 724. Find Pivot Index

給定一整數陣列，請找出 `pivot index`，若查無則回傳 -1。

> pivot index：陣列中的 index，代表左邊元素數值的加總 = 右邊元素數值的加總

規則：

1. index = 0：左方無任一元素，因此加總為 0
2. index = len(array) -1：右方無任一元素，因此加總也是 0
----
[724. Find Pivot Index](https://leetcode.com/problems/find-pivot-index/description)
