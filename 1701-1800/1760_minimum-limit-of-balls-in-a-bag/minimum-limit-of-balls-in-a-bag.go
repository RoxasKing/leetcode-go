package main

// Tags:
// Binary Search

func minimumSize(nums []int, maxOperations int) int {
	l, r := 1, int(1e9)
	for l < r {
		cost := (l + r) >> 1
		if needOperations(nums, cost) > maxOperations {
			l = cost + 1
		} else {
			r = cost
		}
	}
	return l
}

func needOperations(nums []int, cost int) int {
	out := 0
	for _, num := range nums {
		out += num / cost
		if num%cost == 0 {
			out--
		}
	}
	return out
}
