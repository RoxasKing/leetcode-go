package main

import "sort"

// Greedy

func halfQuestions(questions []int) int {
	n := len(questions)
	cnt := map[int]int{}
	for _, q := range questions {
		cnt[q]++
	}
	arr := make([]int, 0, len(cnt))
	for num := range cnt {
		arr = append(arr, num)
	}
	sort.Slice(arr, func(i, j int) bool { return cnt[arr[i]] > cnt[arr[j]] })
	out := 0
	sum := 0
	for _, num := range arr {
		out++
		sum += cnt[num]
		if sum >= n>>1 {
			break
		}
	}
	return out
}
