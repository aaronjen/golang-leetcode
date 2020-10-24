package main

/*
 * @lc app=leetcode id=341 lang=golang
 *
 * [341] Flatten Nested List Iterator
 */

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {
	return false
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
	return 0
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
	return nil
}

// @lc code=start

type NestedIterator struct {
	stack   [][]*NestedInteger
	indices []int
	depth   int
	cur     *NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {

	return &NestedIterator{
		stack:   [][]*NestedInteger{nestedList},
		indices: []int{0},
	}
}

func (this *NestedIterator) Next() int {
	v := this.stack[this.depth][this.indices[this.depth]].GetInteger()
	this.indices[this.depth]++
	return v
}

func (this *NestedIterator) HasNext() bool {
	for {
		if this.depth < 0 {
			return false
		}

		ld := this.stack[this.depth]
		ind := this.indices[this.depth]

		if ind >= len(ld) {
			this.popStack()
			continue
		}

		cur := ld[ind]
		if cur.IsInteger() {
			return true
		}

		list := cur.GetList()
		this.indices[this.depth]++
		if len(list) != 0 {
			this.stack = append(this.stack, list)
			this.indices = append(this.indices, 0)
			this.depth++
		}
	}
}

func (it *NestedIterator) popStack() {
	it.stack = it.stack[:len(it.stack)-1]
	it.indices = it.indices[:len(it.indices)-1]
	it.depth--
}

// @lc code=end
