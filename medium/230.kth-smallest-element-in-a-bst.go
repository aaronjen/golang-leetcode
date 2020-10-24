package main

/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
 */

// TreeNode struct
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

func kthSmallestDFS(node *TreeNode, cnt *int, ans *int) {
	if *cnt <= 0 {
		return
	}
	if node.Left != nil {
		kthSmallestDFS(node.Left, cnt, ans)
	}
	*cnt--
	if *cnt == 0 {
		*ans = node.Val
		return
	}
	if node.Right != nil {
		kthSmallestDFS(node.Right, cnt, ans)
	}

}
func kthSmallest(root *TreeNode, k int) int {
	cnt := k
	ans := 0
	kthSmallestDFS(root, &cnt, &ans)

	return ans
}

// @lc code=end
