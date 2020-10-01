package main

import (
	"math"
)

/*
 * @lc app=leetcode id=530 lang=golang
 *
 * [530] Minimum Absolute Difference in BST
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

func getMinimumDifferenceTraversal(node *TreeNode, prev *int, min *int) {
	if node.Left != nil {
		getMinimumDifferenceTraversal(node.Left, prev, min)
	}

	if *prev == -1 {
		*prev = node.Val
	} else {
		val := node.Val - *prev

		if val < *min {
			*min = val
		}

		*prev = node.Val
	}

	if node.Right != nil {
		getMinimumDifferenceTraversal(node.Right, prev, min)
	}
}

func getMinimumDifference(root *TreeNode) int {
	prev := -1

	min := math.MaxInt32
	getMinimumDifferenceTraversal(root, &prev, &min)

	return min
}

// @lc code=end
