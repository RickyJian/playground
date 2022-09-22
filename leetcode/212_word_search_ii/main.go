package main

import (
	"fmt"
)

func main() {
	board := [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}
	words := []string{"oath", "pea", "eat", "rain"}
	fmt.Println(findWords(board, words))
}

func findWords(board [][]byte, words []string) []string {
	trie := new(node)
	for _, word := range words {
		// 應用 trie 的特性將單字分割成各個英文字母，
		// 待後續走訪 board 時判斷是否能組成單字，
		// 並且減少走訪重複前綴的 word
		trie.insert(word)
	}

	result := make([]string, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(board, i, j, trie, &result)
		}
	}
	return result
}

const wordBeenUsed = '.'

func dfs(board [][]byte, i, j int, trie *node, result *[]string) {
	letter := board[i][j]
	fmt.Printf("i: %d, j: %d, letter: %s\n", i, j, string(letter))
	if letter == wordBeenUsed || trie.children[letter-'a'] == nil {
		// 若遇到以使用過的 cell 或是在 trie 中查無節點，
		// 代表所走訪的 cell 無法組合成單字
		return
	}
	n := trie.children[letter-'a']
	if n.word != "" {
		// 若查找到單字將該單字寫入 result 中，
		*result = append(*result, n.word)
		// 將查找到的單字清除，防止後續找到相同組合
		n.word = ""
		// 有可能會有同時存在前掇綴字相同的單字在此不須回傳
	}

	// 將走訪過的 cell 更成特殊符號，若後續遇到此符號則代表已到盡頭即可跳出 dfs
	board[i][j] = wordBeenUsed
	if i > 0 {
		// 向上找
		dfs(board, i-1, j, n, result)
	}
	if j > 0 {
		// 向前找
		dfs(board, i, j-1, n, result)
	}
	// -1: 最後一個節點不須走訪
	if i < len(board)-1 {
		// 向下找
		dfs(board, i+1, j, n, result)
	}
	// -1: 最後一個節點不須走訪
	if j < len(board[0])-1 {
		// 向後找
		dfs(board, i, j+1, n, result)
	}
	// 將走訪過的 cell 復原
	board[i][j] = letter
}

const alphabetCount = 26

type node struct {
	children [alphabetCount]*node
	word     string
}

func (n *node) insert(word string) {
	for _, w := range word {
		idx := w - 'a'
		if n.children[idx] == nil {
			n.children[idx] = new(node)
		}
		n = n.children[idx]
	}
	n.word = word
}
