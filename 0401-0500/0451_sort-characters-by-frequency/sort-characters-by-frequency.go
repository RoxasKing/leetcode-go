package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Hash
// Sort

func frequencySort(s string) string {
	freq := make([][2]int, 128)
	for i := range freq {
		freq[i][0] = i
	}
	for i := range s {
		freq[s[i]][1]++
	}
	sort.Slice(freq, func(i, j int) bool {
		return freq[i][1] > freq[j][1] || freq[i][1] == freq[j][1] && freq[i][0] < freq[j][0]
	})
	chs := make([]byte, 0, len(s))
	for i := range freq {
		if freq[i][1] == 0 {
			break
		}
		for ; freq[i][1] > 0; freq[i][1]-- {
			chs = append(chs, byte(freq[i][0]))
		}
	}
	return string(chs)
}
