package main

func maxTurbulenceSize(arr []int) int {
	out := 1
	count := 1
	n := len(arr)
	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] {
			if i-2 < 0 || arr[i-2] > arr[i-1] {
				count++
			} else {
				count = 2
			}
		} else if arr[i] < arr[i-1] {
			if i-2 < 0 || arr[i-2] < arr[i-1] {
				count++
			} else {
				count = 2
			}
		} else {
			count = 1
		}
		out = Max(out, count)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
