package main

import "sort"

// Hash
// Sort

func frequencySort(s string) string {
	cnt := [128]int{}
	arr := make([]byte, 128)
	for i := range arr {
		arr[i] = byte(i)
	}
	for i := range s {
		cnt[s[i]]++
	}
	sort.Slice(arr, func(i, j int) bool {
		if cnt[arr[i]] != cnt[arr[j]] {
			return cnt[arr[i]] > cnt[arr[j]]
		}
		return arr[i] < arr[j]
	})
	out := make([]byte, 0, len(s))
	for _, c := range arr {
		for ; cnt[c] > 0; cnt[c]-- {
			out = append(out, c)
		}
	}
	return string(out)
}
