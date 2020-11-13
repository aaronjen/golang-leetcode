package main

import "fmt"

/*
 * @lc app=leetcode id=652 lang=golang
 *
 * [652] Find Duplicate Subtrees
 */

// TreeNode Definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func findDuplicateSubtreesSolve(node *TreeNode, hash map[string]int, res *[]*TreeNode) string {
	if node == nil {
		return "#"
	}
	s := fmt.Sprintf("%v,%s,%s", node.Val, findDuplicateSubtreesSolve(node.Left, hash, res), findDuplicateSubtreesSolve(node.Right, hash, res))

	if hash[s] == 1 {
		*res = append(*res, node)
	}
	hash[s]++

	return s
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	res := []*TreeNode{}
	hash := map[string]int{}
	findDuplicateSubtreesSolve(root, hash, &res)
	return res
}

// @lc code=end
