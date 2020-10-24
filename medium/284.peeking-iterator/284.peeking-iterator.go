package main

/*
 * @lc app=leetcode id=284 lang=golang
 *
 * [284] Peeking Iterator
 */

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return false
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	return 0
}

// @lc code=start

type PeekingIterator struct {
	iter    *Iterator
	current int
	pending bool
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		iter: iter,
	}
}

func (this *PeekingIterator) hasNext() bool {
	if this.pending == false {
		return this.iter.hasNext()
	}

	return true
}

func (this *PeekingIterator) next() int {
	if this.pending {
		this.pending = false
		return this.current
	}

	return this.iter.next()
}

func (this *PeekingIterator) peek() int {
	if this.pending {
		return this.current
	}

	this.pending = true
	this.current = this.iter.next()
	return this.current
}

// @lc code=end
