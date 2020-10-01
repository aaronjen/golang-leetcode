package main

/*
 * @lc app=leetcode id=572 lang=golang
 *
 * [572] Subtree of Another Tree
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

func isSubtreeSame(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	if s.Val != t.Val {
		return false
	}

	return isSubtreeSame(s.Left, t.Left) && isSubtreeSame(s.Right, t.Right)
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	if s.Val == t.Val && isSubtreeSame(s.Left, t.Left) && isSubtreeSame(s.Right, t.Right) {
		return true
	}

	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

// @lc code=end
