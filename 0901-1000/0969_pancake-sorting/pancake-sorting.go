package main

// Difficulty:
// Medium

func pancakeSort(arr []int) []int {
	out := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != i+1 {
			j := i - 1
			for ; arr[j] != i+1; j-- {
			}
			reverse(arr[:j+1])
			reverse(arr[:i+1])
			out = append(out, j+1, i+1)
		}
	}
	return out
}

func reverse(arr []int) {
	for l, r := 0, len(arr)-1; l < r; l, r = l+1, r-1 {
		arr[l], arr[r] = arr[r], arr[l]
	}
}
