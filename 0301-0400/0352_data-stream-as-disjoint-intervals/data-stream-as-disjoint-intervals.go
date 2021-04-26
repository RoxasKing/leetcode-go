package main

import "sort"

/*
  Given a data stream input of non-negative integers a1, a2, ..., an, summarize the numbers seen so far as a list of disjoint intervals.

  Implement the SummaryRanges class:
    1. SummaryRanges() Initializes the object with an empty stream.
    2. void addNum(int val) Adds the integer val to the stream.
    3. int[][] getIntervals() Returns a summary of the integers in the stream currently as a list of disjoint intervals [starti, endi].

  Example 1:
    Input
      ["SummaryRanges", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals"]
      [[], [1], [], [3], [], [7], [], [2], [], [6], []]
    Output
      [null, null, [[1, 1]], null, [[1, 1], [3, 3]], null, [[1, 1], [3, 3], [7, 7]], null, [[1, 3], [7, 7]], null, [[1, 3], [6, 7]]]
    Explanation
      SummaryRanges summaryRanges = new SummaryRanges();
      summaryRanges.addNum(1);      // arr = [1]
      summaryRanges.getIntervals(); // return [[1, 1]]
      summaryRanges.addNum(3);      // arr = [1, 3]
      summaryRanges.getIntervals(); // return [[1, 1], [3, 3]]
      summaryRanges.addNum(7);      // arr = [1, 3, 7]
      summaryRanges.getIntervals(); // return [[1, 1], [3, 3], [7, 7]]
      summaryRanges.addNum(2);      // arr = [1, 2, 3, 7]
      summaryRanges.getIntervals(); // return [[1, 3], [7, 7]]
      summaryRanges.addNum(6);      // arr = [1, 2, 3, 6, 7]
      summaryRanges.getIntervals(); // return [[1, 3], [6, 7]]

  Constraints:
    1. 0 <= val <= 10^4
    2. At most 3 * 10^4 calls will be made to addNum and getIntervals.

  Follow up: What if there are lots of merges and the number of disjoint intervals is small compared to the size of the data stream?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/data-stream-as-disjoint-intervals
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
