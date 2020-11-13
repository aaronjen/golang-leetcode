package main

import "container/list"

/*
 * @lc app=leetcode id=729 lang=golang
 *
 * [729] My Calendar I
 */

// @lc code=start

// MyCalendar struct
type MyCalendar struct {
	events *list.List
}

// Constructor func
func Constructor() MyCalendar {
	return MyCalendar{
		events: list.New(),
	}
}

// Book func
func (cal *MyCalendar) Book(start int, end int) bool {
	events := cal.events
	ev := events.Front()
	for ev != nil {
		elem := ev.Value.([]int)
		eStart, eEnd := elem[0], elem[1]

		if start < eEnd && eStart < end {
			return false
		}
		if eStart < start && eEnd < end {
			ev = ev.Next()
		} else {
			events.InsertBefore([]int{start, end}, ev)
			return true
		}
	}
	events.PushBack([]int{start, end})

	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
// @lc code=end
