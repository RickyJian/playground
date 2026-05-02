---
name: leetcode-gen
description: 根據 LeetCode 題目網址，在 /Users/ricky/github/playground/leetcode/ 底下產生對應的題目資料夾與所有標準檔案（main.go、main_test.go、README.md、go.mod）。當使用者提供 LeetCode 網址並想要生成題目骨架時使用。
disable-model-invocation: true
---

# leetcode-gen

根據 LeetCode 題目網址，自動產生完整的題目資料夾結構。

## 工作流程

### 步驟 1：解析網址

從網址提取資訊：
- 網址格式：`https://leetcode.com/problems/{slug}/`
- **slug**：直接從網址路徑取得（例如 `two-sum`），並將 `-` 替換為 `_` 作為蛇形命名（例如 `two_sum`）

### 步驟 2：抓取題目內容

使用 WebFetch 工具抓取 LeetCode 題目頁面，取得：
- 題目編號（number）
- 題目英文標題（title）
- 題目描述（description）— 翻譯成繁體中文
- 解法函式簽章（function signature）— 從頁面的 code editor 或描述中取得 Go 語言版本

如果 WebFetch 無法取得內容，改抓 `https://leetcode.com/problems/{slug}/description/`。

### 步驟 3：推導命名

| 項目 | 規則 | 範例 |
|------|------|------|
| 資料夾名稱 | `{number}_{slug_underscore}` | `1_two_sum` |
| go.mod module | 將 URL slug 的 `-` 換成 `_`，**不加題號前綴** | `two_sum` |
| 解法函式名 | LeetCode 官方 camelCase | `twoSum` |
| 多版本後綴 | 追加 V2、V3 | `twoSumV2` |

**go.mod module name 命名規則（snake_case，不含題號）：**

| 資料夾 | module name |
|--------|-------------|
| `2099_find_subsequence_of_length_k_with_the_largest_sum` | `find_subsequence_of_length_k_with_the_largest_sum` |
| `2390_removing_stars_from_a_string` | `removing_stars_from_a_string` |
| `219_contains_duplicate_ii` | `contains_duplicate_ii` |
| `1_two_sum` | `two_sum` |

### 步驟 4：建立目錄與檔案

根目錄：`$LEETCODE_HOME/{folder}/`（從環境變數 `LEETCODE_HOME` 取得，若未設定則請先詢問使用者）

依序建立以下四個檔案：

---

#### `main.go`

```go
package main

func main() {
	// do nothing
}

func {functionName}({params}) {returnType} {
	panic("not implemented")
}
```

- `{params}` 與 `{returnType}` 從 LeetCode Go 版函式簽章取得
- 若題目有輔助型別（如 `TreeNode`、`ListNode`），在函式下方定義

---

#### `main_test.go`

```go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test{FunctionName}(t *testing.T) {
	tests := []struct {
		name     string
		// 輸入欄位（對應函式參數）
		expected // 輸出型別
	}{
		{
			name:     "範例 1",
			// 填入 LeetCode 範例測資
		},
		{
			name:     "範例 2",
			// 填入 LeetCode 範例測資
		},
		{
			name:     "邊界條件",
			// 填入邊界或特殊狀況
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := {functionName}(/* tt.欄位 */)
			assert.Equal(t, tt.expected, result)
		})
	}
}
```

規則：
- 測試案例名稱使用**繁體中文**
- 從 LeetCode 範例直接填入測資（至少 2 個範例 + 1 個邊界條件）
- 函式名首字母大寫作為 `Test{FunctionName}`

---

#### `README.md`

```markdown
# {number}. {中文題目名稱}

## 問題描述

{繁體中文題目描述，含範例說明}

## 限制條件

{constraints，保留數學格式}

[{number}. {英文題目名稱}]({leetcode_url})
```

---

#### `go.mod` 與 `go.sum`（透過指令產生）

module name = URL slug 去掉題號、`-` 換成 `_`（全小寫 snake_case）。

在題目資料夾內執行：

```bash
cd $LEETCODE_HOME/{folder}
go mod init {slug_underscore}
go get github.com/stretchr/testify@latest
go mod tidy
```

執行後會自動產生 `go.mod` 與 `go.sum`。`go.sum` 不需手動建立。

**範例（2099 題）：**
```
module find_subsequence_of_length_k_with_the_largest_sum

go 1.x.x

require github.com/stretchr/testify vX.X.X
...
```

---

## 實際範例

**輸入 URL**：`https://leetcode.com/problems/two-sum/`

產生資料夾：`1_two_sum/`

**main.go**：
```go
package main

func main() {
	// do nothing
}

func twoSum(nums []int, target int) []int {
	panic("not implemented")
}
```

**main_test.go**：
```go
package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{name: "範例 1", nums: []int{2, 7, 11, 15}, target: 9, expected: []int{0, 1}},
		{name: "範例 2", nums: []int{3, 2, 4}, target: 6, expected: []int{1, 2}},
		{name: "相同元素", nums: []int{3, 3}, target: 6, expected: []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, twoSum(tt.nums, tt.target))
		})
	}
}
```

**README.md**：
```markdown
# 1. 兩數之和

## 問題描述

給定一個整數陣列 `nums` 和一個整數 `target`，找出陣列中兩個數相加等於 `target` 的索引。

## 限制條件

- `2 <= nums.length <= 10^4`
- `-10^9 <= nums[i] <= 10^9`
- 只會存在唯一解

[1. Two Sum](https://leetcode.com/problems/two-sum/)
```

**go.mod**（執行 `go mod init two_sum && go get github.com/stretchr/testify@latest && go mod tidy` 後產生）：
```
module two_sum

go 1.x.x

require github.com/stretchr/testify vX.X.X

require (
	github.com/davecgh/go-spew vX.X.X // indirect
	github.com/pmezard/go-difflib vX.X.X // indirect
	gopkg.in/yaml.v3 vX.X.X // indirect
)
```

---

## 注意事項

- 若題目涉及 `TreeNode` / `ListNode`，在 `main.go` 末尾加上標準定義（與 LeetCode 一致）
- 若 LeetCode 頁面無法抓取，根據網址 slug 推斷資訊，並在 README 中標示「待補充」
- `go.sum` 由 `go get` 自動產生，不需手動建立
- 每個題目為**獨立 Go module**，不共用 go.mod
