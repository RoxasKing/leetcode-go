package main

// Tags:
// Greedy Algorithm

func isCovered(ranges [][]int, left int, right int) bool {
	for i := range ranges {
		l, r := ranges[i][0], ranges[i][1]
		if l <= left && left <= r {
			left = r + 1
		}
		if l <= right && right <= r {
			right = l - 1
		}
		if left > right {
			break
		}
	}
	return left > right
}
