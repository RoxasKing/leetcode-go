package main

// Difficulty:
// Hard

// Tags:
// Binary Search

func splitArray(nums []int, m int) int {
	l, r := 0, 0
	for _, num := range nums {
		if l < num {
			l = num
		}
		r += num
	}
	for l < r {
		mid := l + (r-l)>>1
		cnt, sum := 1, 0
		for _, num := range nums {
			if sum+num > mid {
				cnt++
				sum = num
			} else {
				sum += num
			}
		}
		if cnt > m {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
