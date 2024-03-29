package main

import "sort"

// Difficulty:
// Hard

// Tags:
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
