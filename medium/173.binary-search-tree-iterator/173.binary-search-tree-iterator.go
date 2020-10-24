package main

/*
 * @lc app=leetcode id=173 lang=golang
 *
 * [173] Binary Search Tree Iterator
 */

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

// BSTIterator struct
type BSTIterator struct {
	stack []*TreeNode
}

// Constructor struct
func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{}
	it.pushAll(root)

	return it
}

func (it *BSTIterator) pushAll(node *TreeNode) {
	for node != nil {
		it.stack = append(it.stack, node)
		node = node.Left
	}
}

// Next func
/** @return the next smallest number */
func (it *BSTIterator) Next() int {
	n := len(it.stack)
	node := it.stack[n-1]
	it.stack = it.stack[:n-1]
	it.pushAll(node.Right)

	return node.Val
}

// HasNext func
/** @return whether we have a next smallest number */
func (it *BSTIterator) HasNext() bool {
	return len(it.stack) != 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end
