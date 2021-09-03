package main

import "math/rand"

// Tags:
// Quick Select

func smallestK(arr []int, k int) []int {
	for l, r := 0, len(arr)-1; l < r; {
		x := l + rand.Intn(r+1-l)
		arr[x], arr[r] = arr[r], arr[x]
		i, j := l-1, r
		for {
			for i++; i < r && arr[i] < arr[r]; i++ {
			}
			for j--; j > l && arr[j] > arr[r]; j-- {
			}
			if i >= j {
				break
			}
			arr[i], arr[j] = arr[j], arr[i]
		}
		arr[i], arr[r] = arr[r], arr[i]
		if i < k {
			l = i + 1
		} else if i > k {
			r = i - 1
		} else {
			break
		}
	}
	return arr[:k]
}
