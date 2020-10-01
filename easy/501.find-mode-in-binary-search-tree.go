package main

/*
 * @lc app=leetcode id=501 lang=golang
 *
 * [501] Find Mode in Binary Search Tree
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findModeDFS(node *TreeNode, m map[int]int) {
	if node.Left != nil {
		findModeDFS(node.Left, m)
	}

	if node.Right != nil {
		findModeDFS(node.Right, m)
	}

	m[node.Val]++
}

func findMode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	h := map[int]int{}

	findModeDFS(root, h)

	max := 0
	ret := []int{}
	for mode, val := range h {
		if val > max {
			ret = []int{mode}
			max = val
		} else if val == max {
			ret = append(ret, mode)
		}
	}

	return ret
}

// @lc code=end
