package main

/*
 * @lc app=leetcode id=687 lang=golang
 *
 * [687] Longest Univalue Path
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func longestUnivaluePathSolve(node *TreeNode, value int, max *int) int {
	if node == nil {
		return 0
	}

	l := longestUnivaluePathSolve(node.Left, node.Val, max)
	r := longestUnivaluePathSolve(node.Right, node.Val, max)

	if l+r > *max {
		*max = l + r
	}

	if node.Val == value {
		if l > r {
			return l + 1
		}
		return r + 1
	}

	return 0
}

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := 0

	longestUnivaluePathSolve(root, root.Val, &max)

	return max
}

// @lc code=end
