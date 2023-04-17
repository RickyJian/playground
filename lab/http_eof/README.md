# HTTP Client EOF 問題

## 問題

當 client 對 server 發起 request 時，有時會在 client 端遇到 `EOF` 的錯誤，但 server 端卻沒有顯示任何的錯誤。

### EOF 原因

由於 HTTP response write timeout 導致後續 client 端  

----

## 參考資料：

* https://ieftimov.com/posts/make-resilient-golang-net-http-servers-using-timeouts-deadlines-context-cancellation/
* https://juejin.cn/post/6969184203887869983