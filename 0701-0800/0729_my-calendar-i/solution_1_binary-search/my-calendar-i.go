package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Binary Search

type MyCalendar struct {
	a [][]int
	n int
}

func Constructor() MyCalendar {
	return MyCalendar{make([][]int, 1000), 0}
}

func (this *MyCalendar) Book(start int, end int) bool {
	l, r := start, end-1
	i := sort.Search(this.n, func(i int) bool { return this.a[i][0] > l })
	if i > 0 && this.a[i-1][1] >= l || i < this.n && this.a[i][0] <= r {
		return false
	}
	copy(this.a[i+1:], this.a[i:])
	this.a[i] = []int{l, r}
	this.n++
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
