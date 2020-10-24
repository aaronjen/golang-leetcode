package main

import (
	"container/heap"
)

/*
 * @lc app=leetcode id=373 lang=golang
 *
 * [373] Find K Pairs with Smallest Sums
 */

// @lc code=start
type item struct {
	i1 int
	i2 int
	v  int
}

type minHeap []*item

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].v < h[j].v }
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

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return [][]int{}
	}

	res := [][]int{}

	h := minHeap{}
	cnt := 0
	gi1 := 0
	for ; gi1 < len(nums1); gi1++ {
		if cnt >= k {
			break
		}
		h = append(h, &item{
			i1: gi1,
			i2: 0,
			v:  nums1[gi1] + nums2[0],
		})
		cnt++
	}

	heap.Init(&h)

	for i := 0; i < k; i++ {
		if len(h) == 0 {
			break
		}
		it := heap.Pop(&h).(*item)
		i1 := it.i1
		i2 := it.i2

		res = append(res, []int{nums1[i1], nums2[i2]})
		if i2+1 < len(nums2) {
			heap.Push(&h, &item{
				i1,
				i2 + 1,
				nums1[i1] + nums2[i2+1],
			})
		} else {
			gi1++
			if gi1 >= len(nums1) {
				continue
			}
			heap.Push(&h, &item{
				gi1,
				i2,
				nums1[gi1] + nums2[i2],
			})
		}

	}

	return res
}

// @lc code=end

// "[1, 1, 2]\n[1, 2, 3]\n10"
