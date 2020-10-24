package main

import "math/rand"

/*
 * @lc app=leetcode id=380 lang=golang
 *
 * [380] Insert Delete GetRandom O(1)
 */

// @lc code=start

// RandomizedSet struct
type RandomizedSet struct {
	table map[int]int
	arr   []int
}

// Constructor func
/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		table: map[int]int{},
	}
}

// Insert func
/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.table[val]; ok {
		return false
	}
	n := len(s.arr)
	s.table[val] = n
	s.arr = append(s.arr, val)

	return true
}

// Remove func
/** Removes a value from the set. Returns true if the set contained the specified element. */
func (s *RandomizedSet) Remove(val int) bool {
	if ind, ok := s.table[val]; ok {
		n := len(s.arr)
		last := s.arr[n-1]
		s.table[last] = ind
		s.arr[ind] = last
		s.arr = s.arr[:n-1]
		delete(s.table, val)

		return true
	}

	return false
}

// GetRandom func
/** Get a random element from the set. */
func (s *RandomizedSet) GetRandom() int {
	ind := rand.Intn(len(s.table))

	return s.arr[ind]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
