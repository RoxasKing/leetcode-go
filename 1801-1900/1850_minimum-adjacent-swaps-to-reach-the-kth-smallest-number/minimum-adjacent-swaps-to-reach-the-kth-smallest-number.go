package main

import (
	"sort"
)

func getMinSwaps(num string, k int) int {
	chs := []byte(num)
	for ; k > 0; k-- {
		nextPermutation(chs)
	}

	out := 0
	for i := range chs {
		if chs[i] != num[i] {
			j := i + 1
			for chs[j] != num[i] {
				j++
			}
			for ; j > i; j-- {
				chs[j], chs[j-1] = chs[j-1], chs[j]
				out++
			}
		}
	}
	return out
}

func nextPermutation(chs []byte) {
	n := len(chs)
	i := n - 1
	for i-1 >= 0 && chs[i-1] >= chs[i] {
		i--
	}
	reverse(chs[i:])

	j := sort.Search(n-i, func(j int) bool { return chs[j+i] > chs[i-1] }) + i
	chs[i-1], chs[j] = chs[j], chs[i-1]
}

func reverse(chs []byte) {
	n := len(chs)
	for i := 0; i < n>>1; i++ {
		chs[i], chs[n-1-i] = chs[n-1-i], chs[i]
	}
}
