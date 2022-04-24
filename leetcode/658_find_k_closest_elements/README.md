# 658. Find K Closest Elements

給定一個排序過的陣列，以及 k 和 x 正整數，請找出最靠近 x 的 k 個值並以降冪排序回傳陣列。

> 有 a 與 b 倆值，
>   1. a 值較 b 值靠近 x (|a - x| < |b - x|)
>   2. a 與 b 離 x 值相同，但 a < b (|a - x| == |b - x| and a < b)
> 
> a 值為解答

[Find K Closest Elements](https://leetcode.com/problems/find-k-closest-elements/)