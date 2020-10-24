package main

/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */

// @lc code=start

// TrieNode struct
type TrieNode struct {
	isWord   bool
	children [26]*TrieNode
}

// Trie struct
type Trie struct {
	root *TrieNode
}

// Constructor func
/** Initialize your data structure here. */
func Constructor() Trie {
	root := new(TrieNode)
	root.isWord = false

	return Trie{
		root: root,
	}
}

// Insert func
/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	nd := t.root

	for i := 0; i < len(word); i++ {
		b := word[i]

		ind := int(b - 'a')

		if nd.children[ind] == nil {
			nd.children[ind] = new(TrieNode)
		}
		nd = nd.children[ind]
	}
	nd.isWord = true
}

// Search func
/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	nd := t.root

	for i := 0; i < len(word); i++ {
		b := word[i]

		ind := int(b - 'a')

		if nd.children[ind] == nil {
			return false
		}
		nd = nd.children[ind]
	}

	return nd.isWord
}

// StartsWith func
/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	nd := t.root

	for i := 0; i < len(prefix); i++ {
		b := prefix[i]

		ind := int(b - 'a')

		if nd.children[ind] == nil {
			return false
		}

		nd = nd.children[ind]
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
