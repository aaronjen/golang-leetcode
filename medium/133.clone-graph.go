package main

/*
 * @lc app=leetcode id=133 lang=golang
 *
 * [133] Clone Graph
 */

// Node struct
type Node struct {
	Val       int
	Neighbors []*Node
}

// @lc code=start

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	checked := map[int]bool{}
	nodeTable := map[int]*Node{}

	q := []*Node{node}

	for len(q) != 0 {
		n := q[0]
		if !checked[n.Val] {
			checked[n.Val] = true

			cur := nodeTable[n.Val]
			if cur == nil {
				cur = new(Node)
				nodeTable[n.Val] = cur
				cur.Val = n.Val
			}

			for _, neighbor := range n.Neighbors {
				if !checked[neighbor.Val] {
					q = append(q, neighbor)
				}

				nn := nodeTable[neighbor.Val]
				if nn == nil {
					tmp := &Node{
						Val: neighbor.Val,
					}
					nodeTable[neighbor.Val] = tmp
					nn = tmp
				}
				cur.Neighbors = append(cur.Neighbors, nn)
			}

		}

		q = q[1:]
	}

	return nodeTable[node.Val]
}

// @lc code=end
