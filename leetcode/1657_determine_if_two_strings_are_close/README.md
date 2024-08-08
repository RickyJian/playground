# 1657. Determine if Two Strings Are Close

若能通過以下操作，將字串 `word1` 並更成字串 `word2`，則可以將這兩個字串稱作『接近』。

操作：

1. 任意交換兩個字串中字的位置。（例：a**b**cd**e** -> a**e**cd**b**）
2. 將每個出現的字轉換成另一個出現的字（所有相同的字皆須變成另一個相同的。例： **aa**c**abb** -> **bb**c**baa**）

> 不限操作次數

若能將 `word1` 轉換成 `word2` 則回傳 true，反之回傳 false。

-----
[1657. Determine if Two Strings Are Close](https://leetcode.com/problems/determine-if-two-strings-are-close)