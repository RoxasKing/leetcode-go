package main

import (
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

// Tags:
// Quick Sort
func minNumber(nums []int) string {
	n := len(nums)
	strs := make([]string, n)
	for i := range nums {
		strs[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(strs, func(i, j int) bool { return strs[i]+strs[j] < strs[j]+strs[i] })
	return strings.Join(strs, "")
}

func minNumber2(nums []int) string {
	n := len(nums)
	strs := make([]string, n)
	for i := range nums {
		strs[i] = strconv.Itoa(nums[i])
	}
	qSort(strs, 0, n-1)
	return strings.Join(strs, "")
}

func qSort(strs []string, l, r int) {
	if l >= r {
		return
	}
	pivotIdx := l + rand.Intn(r+1-l)
	strs[pivotIdx], strs[r] = strs[r], strs[pivotIdx]
	i := l
	for j := l; j < r; j++ {
		if isSmaller(strs[j], strs[r]) {
			strs[j], strs[i] = strs[i], strs[j]
			i++
		}
	}
	strs[i], strs[r] = strs[r], strs[i]
	qSort(strs, l, i-1)
	qSort(strs, i+1, r)
}

func isSmaller(a, b string) bool {
	a, b = a+b, b+a
	return a < b
}
