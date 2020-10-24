package main

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
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
func binaryTreePathsDFS(node *TreeNode, res *[]string, prefix string) {
	s := prefix + fmt.Sprintf("->%v", node.Val)

	if node.Left != nil {
		binaryTreePathsDFS(node.Left, res, s)
	}
	if node.Right != nil {
		binaryTreePathsDFS(node.Right, res, s)
	}

	if node.Left == nil && node.Right == nil {
		*res = append(*res, s)
	}
}

func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	if root == nil {
		return res
	}

	s := strconv.Itoa(root.Val)

	if root.Left != nil {
		binaryTreePathsDFS(root.Left, &res, s)
	}
	if root.Right != nil {
		binaryTreePathsDFS(root.Right, &res, s)
	}

	if root.Left == nil && root.Right == nil {
		res = append(res, s)
	}

	return res
}

// @lc code=end
