package main

// Tags:
// Binary Search

func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1
	for l+1 < r {
		m := (l + r) >> 1
		if arr[m] > arr[m-1] && arr[m] > arr[m+1] {
			return m
		}
		if arr[m] <= arr[m+1] {
			l = m
		} else {
			r = m
		}
	}
	return -1
}
