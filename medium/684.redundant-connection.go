package main

/*
 * @lc app=leetcode id=684 lang=golang
 *
 * [684] Redundant Connection
 */

// @lc code=start
func findRoot(tree []int, nd int) int {
	if tree[nd] != nd {
		tree[nd] = findRoot(tree, tree[nd])
	}

	return tree[nd]
}

func findRedundantConnection(edges [][]int) []int {
	tree := make([]int, len(edges)+1)
	for i := 0; i < len(tree); i++ {
		tree[i] = i
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		uRoot := findRoot(tree, u)
		vRoot := findRoot(tree, v)

		if uRoot == vRoot {
			return edge
		}

		tree[uRoot] = vRoot
	}

	return []int{}
}

// @lc code=end
