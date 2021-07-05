package main

// Tags:
// Binary Search

func search(arr []int, target int) int {
	n := len(arr)
	l, r := 0, n-1
	for l < r && arr[l] == arr[r] {
		r--
	}
	rotateIdx := bSearchRotateIdx(arr, l, r)
	if rotateIdx == 0 || arr[l] > target {
		return bSearch(arr, rotateIdx, r, target)
	}
	return bSearch(arr, l, rotateIdx, target)
}

func bSearchRotateIdx(arr []int, l, r int) int {
	start := l
	for l < r {
		m := (l + r) >> 1
		if arr[m] > arr[m+1] {
			return m + 1
		}
		if arr[l] <= arr[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	return start
}

func bSearch(arr []int, l, r, target int) int {
	for l < r {
		m := (l + r) >> 1
		if arr[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	if arr[l] == target {
		return l
	}
	return -1
}
