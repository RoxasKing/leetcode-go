package main

// Difficulty:
// Easy

// Tags:
// Hash

func canFormArray(arr []int, pieces [][]int) bool {
	h := [101]int{}
	for i := 1; i <= 100; i++ {
		h[i] = -1
	}
	for i, x := range arr {
		h[x] = i
	}
	for _, a := range pieces {
		i := h[a[0]]
		if i == -1 || len(a) > len(arr)-i {
			return false
		}
		for j, x := range a {
			if arr[i+j] != x {
				return false
			}
		}
	}
	return true
}
