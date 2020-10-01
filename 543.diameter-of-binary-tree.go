package main

/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
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
func diameterOfBinaryTreeDFS(node *TreeNode, max *int) int {
	if node == nil {
		return 0
	}
	leftDepth := diameterOfBinaryTreeDFS(node.Left, max)
	rightDepth := diameterOfBinaryTreeDFS(node.Right, max)

	if leftDepth+rightDepth > *max {
		*max = leftDepth + rightDepth
	}

	if leftDepth > rightDepth {
		return leftDepth + 1
	}

	return rightDepth + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	max := 0
	diameterOfBinaryTreeDFS(root, &max)

	return max
}

// @lc code=end
