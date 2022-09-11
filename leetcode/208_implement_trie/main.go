package main

import (
	"fmt"
)

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
func main() {
	// 	["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
	// [[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
	t := Constructor()
	t.Insert("apple")
	fmt.Println("find: apple", t.Search("apple"))
	fmt.Println("find: app", t.Search("app"))
	fmt.Println("start with: app", t.StartsWith("app"))
	t.Insert("app")
	fmt.Println("find: apple", t.Search("apple"))
}

const (
	alphabetCount = 26
)

type Trie struct {
	children [alphabetCount]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	cur := this
	for _, w := range word {
		idx := w - 'a'
		if cur.children[idx] == nil {
			cur.children[idx] = new(Trie)
		}
		cur = cur.children[idx]
	}
	cur.isEnd = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	for _, w := range word {
		idx := w - 'a'
		if cur.children[idx] == nil {
			return false
		}
		cur = cur.children[idx]
	}
	return cur.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, w := range prefix {
		idx := w - 'a'
		if cur.children[idx] == nil {
			return false
		}
		cur = cur.children[idx]
	}
	return true
}
