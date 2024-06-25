# 443. String Compression

給定一個字元陣列 `chars`，該陣列存在每一組連續重複的字符，稱其為 `group`，

從空字串 `s` 開始，請根據以下規則將其壓縮：

1. 若 `group` 僅有一個字元，直接將該字元加入 `s`
2. 若 `group` 有多個字元，統計次數並將字元級次數加入 `s`

壓縮後的 `s` 不應單獨返回須存入 `chars` 中。

> 長度 >= 10 的 `group` 將被切分並存入 `chars` 中

限制：空間複雜度須為 O(n)
-------
[443. String Compression](https://leetcode.com/problems/string-compression)
