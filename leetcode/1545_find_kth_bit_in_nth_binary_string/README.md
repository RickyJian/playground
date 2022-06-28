# 1545. Find Kth Bit in Nth Binary String

給定 n, k 的正整數，請根據下方方程式找出第 k 的位置的值。

> S1 = "0" <br>
> Sn = Sn - 1 + "1" + reverse(invert(Sn - 1)) for n > 1

範例：

> S1 = "0" <br>
> S2 = "011" <br>
> S3 = "0111001" <br>
> S4 = "011100110110001"

[Find Kth Bit in Nth Binary String](https://leetcode.com/problems/find-kth-bit-in-nth-binary-string/)
