package main

import (
	"math/rand"
)

// Tags:
// Bit Manipulation
// Prefix Sum
// Quick Select

func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	arr := make([]int, 0, m*n)
	f := make([][]int, m+1)
	f[0] = make([]int, n+1)
	for i := 0; i < m; i++ {
		f[i+1] = make([]int, n+1)
		for j := 0; j < n; j++ {
			f[i+1][j+1] = matrix[i][j] ^ f[i][j+1] ^ f[i+1][j] ^ f[i][j]
			arr = append(arr, f[i+1][j+1])
		}
	}

	return quickSelect(arr, m*n-k)
}

func quickSelect(arr []int, k int) int {
	for l, r := 0, len(arr)-1; l < r; {
		pivotIdx := rand.Intn(r+1-l) + l
		arr[pivotIdx], arr[r] = arr[r], arr[pivotIdx]
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
	return arr[k]
}
