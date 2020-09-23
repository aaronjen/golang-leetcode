package main

/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

// @lc code=start
type MinStack struct {
	nums []int
	mins []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.nums = append(this.nums, x)
	if len(this.mins) == 0 || x <= this.mins[len(this.mins)-1] {
		this.mins = append(this.mins, x)
	}
}

func (this *MinStack) Pop() {
	nums := this.nums
	mins := this.mins
	this.nums = nums[:len(nums)-1]
	if nums[len(nums)-1] == mins[len(mins)-1] {
		this.mins = mins[:len(mins)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.nums) == 0 {
		return 0
	}
	return this.nums[len(this.nums)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.mins) == 0 {
		return 0
	}

	return this.mins[len(this.mins)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
