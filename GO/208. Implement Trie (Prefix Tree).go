package main

import "fmt"

type Trie struct {
	children []*Trie
	isEnd bool
}


func Constructor() Trie {
	return Trie{
		children : make([]*Trie, 26),
		isEnd: false,
	}
}


func (this *Trie) Insert(word string)  {
	t := this
	for i := 0; i < len(word); i++ {
		char := word[i] - 'a'
		if t.children[char] == nil {
			t.children[char] = &Trie{
				children : make([]*Trie, 26),
				isEnd: false,
			}
		}
		t = t.children[char]
	}
	t.isEnd = true
}


func (this *Trie) Search(word string) bool {
	t := this
	for i := 0; i < len(word); i++ {
		char := word[i] - 'a'
		if t.children[char] == nil {
			return false
		} else {
			t = t.children[char]
		}
	}
	return t != nil && t.isEnd
}


func (this *Trie) StartsWith(prefix string) bool {
	t := this
	for i := 0; i < len(prefix); i++ {
		char := prefix[i] - 'a'
		if t.children[char] == nil {
			return false
		} else {
			t = t.children[char]
		}
	}
	return t != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	
}