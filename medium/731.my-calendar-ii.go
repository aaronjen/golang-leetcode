package main

import (
	"container/list"
	"fmt"
)

/*
 * @lc app=leetcode id=731 lang=golang
 *
 * [731] My Calendar II
 */

// @lc code=start

// MyCalendarTwo struct
type MyCalendarTwo struct {
	events *list.List
}

// Constructor MyCalendarTwo
func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		events: list.New(),
	}
}

// Book func
func (cal *MyCalendarTwo) Book(start int, end int) bool {
	events := cal.events
	ev := events.Front()
	conflict := 0
	for ev != nil {
		elem := ev.Value.([]int)
		eStart, eEnd := elem[0], elem[1]
		fmt.Println(start, end, eStart, eEnd)
		if start < eEnd && eStart < end {
			if conflict == 1 {
				return false
			}
			conflict++
		} else if eStart > start && eEnd > end {
			if end == eStart {
				elem[0] = start
			} else {
				events.InsertBefore([]int{start, end}, ev)
			}
			return true
		}
		ev = ev.Next()
	}
	events.PushBack([]int{start, end})

	return true
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
// @lc code=end
