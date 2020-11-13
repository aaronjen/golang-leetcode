package main

import (
	"container/heap"
)

/*
 * @lc app=leetcode id=378 lang=golang
 *
 * [378] Kth Smallest Element in a Sorted Matrix
 */

// @lc code=start
type item struct {
	row int
	col int
	val int
}

type minHeap []*item

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*item))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	h := make(minHeap, n)

	for i := 0; i < n; i++ {
		h[i] = &item{
			row: i,
			col: 0,
			val: matrix[i][0],
		}
	}

	heap.Init(&h)

	cnt := 0
	for len(h) != 0 {
		it := heap.Pop(&h).(*item)
		cnt++
		if cnt == k {
			return it.val
		}
		it.col++
		if it.col < n {
			it.val = matrix[it.row][it.col]
			heap.Push(&h, it)
		}
	}

	return 0
}

// @lc code=end
