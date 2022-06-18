package main

// Difficulty:
// Easy

// Difficulty:
// Two Pointers

func duplicateZeros(arr []int) {
	n := len(arr)
	k, c := 0, 0
	for i, x := range arr {
		c++
		if x == 0 {
			c++
		}
		if c >= n {
			k = i
			break
		}
	}
	i := n - 1
	if c > n && arr[k] == 0 {
		arr[i] = arr[k]
		k--
		i--
	}
	for ; k >= 0; k-- {
		arr[i] = arr[k]
		i--
		if arr[k] == 0 {
			arr[i] = arr[k]
			i--
		}
	}
}
