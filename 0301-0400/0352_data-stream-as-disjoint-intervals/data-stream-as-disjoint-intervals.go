package main

import "sort"

// Binary Search
type SummaryRanges struct {
	ranges [][]int
}

/** Initialize your data structure here. */
func Constructor() SummaryRanges {
	return SummaryRanges{}
}

func (this *SummaryRanges) AddNum(val int) {
	idx := sort.Search(len(this.ranges), func(i int) bool {
		return this.ranges[i][0] >= val
	})

	if idx > 0 && val <= this.ranges[idx-1][1] || idx < len(this.ranges) && val == this.ranges[idx][0] {
		return
	}

	if idx > 0 && this.ranges[idx-1][1]+1 == val {
		this.ranges[idx-1][1]++
	} else if idx < len(this.ranges) && val+1 == this.ranges[idx][0] {
		this.ranges[idx][0]--
	} else {
		if idx < len(this.ranges) {
			this.ranges = append(this.ranges, []int{})
			copy(this.ranges[idx+1:], this.ranges[idx:])
			this.ranges[idx] = []int{val, val}
		} else {
			this.ranges = append(this.ranges, []int{val, val})
		}
		return
	}

	if 0 < idx && idx < len(this.ranges) && this.ranges[idx-1][1]+1 == this.ranges[idx][0] {
		this.ranges[idx-1][1] = this.ranges[idx][1]
		if idx+1 < len(this.ranges) {
			copy(this.ranges[idx:], this.ranges[idx+1:])
		}
		this.ranges = this.ranges[:len(this.ranges)-1]
	}
}

func (this *SummaryRanges) GetIntervals() [][]int {
	return this.ranges
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(val);
 * param_2 := obj.GetIntervals();
 */
