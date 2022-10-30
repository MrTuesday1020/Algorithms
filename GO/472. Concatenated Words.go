package main

import (
	"fmt"
	"sort"
)

type Trie struct {
	children []*Trie
	isEnd bool
}

func Constructor() *Trie {
	return &Trie{
		children: make([]*Trie, 26),
		isEnd: false,
	}
}

func (this *Trie) Insert(word string) {
	t := this
	for i := 0; i < len(word); i++ {
		char := word[i] - 'a'
		if t.children[char] == nil {
			t.children[char] = Constructor()
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
		}
		t = t.children[char]
	}
	return t != nil && t.isEnd
}

func findAllConcatenatedWordsInADict(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	
	fmt.Println("sort", words)
			
	var ret []string
	trie := Constructor()
	for i := 0; i < len(words); i++ {
		if check(trie, words[i]) {
			ret = append(ret, words[i])
		}
		trie.Insert(words[i])
	}
	
	return ret
}

func check(trie *Trie, word string) bool {
	fmt.Println("check", word)
	if word == "" {
		return true
	}
	var ret bool
	for i := 1; i < len(word)+1; i++ {
		if trie.Search(word[:i]) {
			ret = ret || check(trie, word[i:])
		} 
	}
	return ret
}

func main() {
	fmt.Println(findAllConcatenatedWordsInADict([]string{"cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"}))
}