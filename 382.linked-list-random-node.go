package main

import "math/rand"

/*
 * @lc app=leetcode id=382 lang=golang
 *
 * [382] Linked List Random Node
 */

// ListNode struct
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

// Solution struct
type Solution struct {
	head *ListNode
}

// Constructor func
/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	return Solution{
		head: head,
	}
}

// GetRandom func
/** Returns a random node's value. */
func (s *Solution) GetRandom() int {
	res := s.head.Val
	nd := s.head
	for i := 1; nd.Next != nil; i++ {
		nd = nd.Next
		if rand.Intn(i+1) == i {
			res = nd.Val
		}
	}

	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
// @lc code=end
