package main

import (
	"sort"
)

/*
  You are given an integer array nums and you have to return a new counts array. The counts array has the property where counts[i] is the number of smaller elements to the right of nums[i].

  Example 1:
    Input: nums = [5,2,6,1]
    Output: [2,1,1,0]
    Explanation:
      To the right of 5 there are 2 smaller elements (2 and 1).
      To the right of 2 there is only 1 smaller element (1).
      To the right of 6 there is 1 smaller element (1).
      To the right of 1 there is 0 smaller element.

  Example 2:
    Input: nums = [-1]
    Output: [0]

  Example 3:
    Input: nums = [-1,-1]
    Output: [0,0]

  Constraints:
    1. 1 <= nums.length <= 10^5
    2. -10^4 <= nums[i] <= 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Fenwick Tree
func countSmaller(nums []int) []int {
	n := len(nums)
	out := make([]int, 0, n)
	ft := NewFenwickTree()
	ft.Discretization(nums, n)

	for i := n - 1; i >= 0; i-- {
		id := ft.GetId(nums[i])
		out = append(out, ft.Query(id-1))
		ft.Update(id)
	}

	for i := 0; i < n>>1; i++ {
		out[i], out[n-1-i] = out[n-1-i], out[i]
	}

	return out
}

type FenwickTree struct {
	a, c []int
}

func NewFenwickTree() FenwickTree {
	return FenwickTree{}
}

func (ft *FenwickTree) Discretization(nums []int, n int) {
	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}
	ft.a = make([]int, 0, n)
	for num := range set {
		ft.a = append(ft.a, num)
	}
	sort.Ints(ft.a)
	ft.c = make([]int, n)
}

func (ft *FenwickTree) GetId(x int) int {
	return sort.SearchInts(ft.a, x) + 1
}

func (ft *FenwickTree) Update(id int) {
	for id < len(ft.c) {
		ft.c[id]++
		id += lowBit(id)
	}
}

func (ft *FenwickTree) Query(id int) int {
	out := 0
	for id > 0 {
		out += ft.c[id]
		id -= lowBit(id)
	}
	return out
}

func lowBit(x int) int {
	return x & (-x)
}
