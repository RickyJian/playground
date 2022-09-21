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
	if letter == wordBeenUsed || trie.children[letter-'a'] == nil {
		return
	}
	n := trie.children[letter-'a']
	if n.word != "" {
		*result = append(*result, n.word)
		// remove used word
		n.word = ""
	}

	board[i][j] = wordBeenUsed
	if i > 0 {
		// 向後找
		dfs(board, i-1, j, n, result)
	}
	if j > 0 {
		// 向上找
		dfs(board, i, j-1, n, result)
	}
	if i < len(board)-1 {
		// 向前找
		dfs(board, i+1, j, n, result)
	}
	if j < len(board[0])-1 {
		// 向下找
		dfs(board, i, j+1, n, result)
	}
	// 結束須將值回復
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
