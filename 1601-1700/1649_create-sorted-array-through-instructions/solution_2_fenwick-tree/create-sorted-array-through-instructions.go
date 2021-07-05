package main

import "sort"

// Fenwick Tree

func createSortedArray(instructions []int) int {
	set := map[int]struct{}{}
	arr := []int{}
	for _, num := range instructions {
		if _, ok := set[num]; !ok {
			set[num] = struct{}{}
			arr = append(arr, num)
		}
	}
	sort.Ints(arr)

	n := len(arr)
	ft := make([]int, n+1)
	out := 0

	for i, num := range instructions {
		idx := sort.SearchInts(arr, num) + 1
		if 1 < idx && idx < n {
			out += Min(query(ft, idx-1), i-query(ft, idx))
		}
		add(ft, idx)
	}

	return out % (1e9 + 7)
}

func add(ft []int, idx int) {
	for idx < len(ft) {
		ft[idx]++
		idx += lowBit(idx)
	}
}

func query(ft []int, idx int) int {
	out := 0
	for idx > 0 {
		out += ft[idx]
		idx -= lowBit(idx)
	}
	return out
}

func lowBit(x int) int {
	return x & (-x)
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
