package main

import "sort"

// Tags:
// Hash
// Sorting

func groupAnagrams(strs []string) [][]string {
	dict := map[string][]string{}
	for _, str := range strs {
		arr := []byte(str)
		sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
		dict[string(arr)] = append(dict[string(arr)], str)
	}
	out := [][]string{}
	for _, strs := range dict {
		out = append(out, strs)
	}
	return out
}
