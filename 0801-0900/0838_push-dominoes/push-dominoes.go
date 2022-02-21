package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func pushDominoes(dominoes string) string {
	n := len(dominoes)
	fl := make([]int, n)
	fr := make([]int, n)
	for i := 0; i < n; i++ {
		fl[i] = 1e5
		fr[i] = 1e5
	}
	for i := 0; i < n; i++ {
		if dominoes[i] == 'R' {
			fr[i] = 0
		} else if dominoes[i] == '.' && i > 0 {
			fr[i] = Min(fr[i], fr[i-1]+1)
		}
	}
	for i := n - 1; i >= 0; i-- {
		if dominoes[i] == 'L' {
			fl[i] = 0
		} else if dominoes[i] == '.' && i < n-1 {
			fl[i] = Min(fl[i], fl[i+1]+1)
		}
	}
	arr := make([]byte, n)
	for i := 0; i < n; i++ {
		if fl[i] < fr[i] {
			arr[i] = 'L'
		} else if fl[i] > fr[i] {
			arr[i] = 'R'
		} else {
			arr[i] = '.'
		}
	}
	return string(arr)
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
