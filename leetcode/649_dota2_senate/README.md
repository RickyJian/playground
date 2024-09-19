# 649. Dota2 Senate

在Dota2的世界中，有兩個政黨:`Radiant` 和 `Dire`。

Dota2的參議院由來自這兩個政黨的參議員組成。

現在參議院想要對Dota2遊戲做出一項改變，。在每個回合中每位參議員可以行使以下兩項權利之一:

* 禁止一位參議員的權利:一位參議員可以使另一位參議員在本回合及所有後續回合中失去所有權利。
* 宣布勝利:如果這位參議員發現仍有投票權的參議員都來自同一政黨,他可以宣布勝利並決定對遊戲的改變。

給定一個字串 `senate` 代表每位參議員所屬的政黨 `R` 和 `D` 分別代表`Radiant` 和 `Dire`。

如果有n位參議員,則給定字串的長度為n。

每次從第一位參議員開始，按給定順序進行到最後一位參議員將持續到投票結束。

失去權利的參議員會喪失後續回合的全力。

假設每位參議員都足夠聰明，會為自己的政黨採取最佳策略。請預測哪個政黨最終會宣布勝利。

輸出為 `Radiant`或 `Dire`。

----

[649. Dota2 Senate](https://leetcode.com/problems/dota2-senate/description)