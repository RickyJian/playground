# 241. Different Ways to Add Parentheses

給定一個算式(例：`2-1-1`)，用括弧群組各數值及運算元，並計算出所有結果。

範例1：

```
Input: expression = "2-1-1"
Output: [0,2]
Explanation:
((2-1)-1) = 0 
(2-(1-1)) = 2
```

範例2：

```
Input: expression = "2*3-4*5"
Output: [-34,-14,-10,-10,10]
Explanation:
(2*(3-(4*5))) = -34 
((2*3)-(4*5)) = -14 
((2*(3-4))*5) = -10 
(2*((3-4)*5)) = -10 
(((2*3)-4)*5) = 10
```

補充說明：

1. 1 <= expression.length <= 20
2. expression consists of digits and the operator '+', '-', and '*'.
3. All the integer values in the input expression are in the range [0, 99].

[Different Ways to Add Parentheses](https://leetcode.com/problems/different-ways-to-add-parentheses/)