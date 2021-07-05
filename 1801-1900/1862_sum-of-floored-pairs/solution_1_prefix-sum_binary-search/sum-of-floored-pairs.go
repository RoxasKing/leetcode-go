package main

import "sort"

// Prefix Sum + Binary Search

func sumOfFlooredPairs(nums []int) int {
	freq := map[int]int{}
	for _, num := range nums {
		freq[num]++
	}
	arr := make([]int, 0, len(freq))
	for num := range freq {
		arr = append(arr, num)
	}
	sort.Ints(arr)

	n := len(arr)
	pFreq := make([]int, n+1)
	for i := 0; i < n; i++ {
		pFreq[i+1] = pFreq[i] + freq[arr[i]]
	}

	out := 0
	for i, num := range arr {
		for base := num; i < n; base += num {
			j := sort.Search(n-i, func(j int) bool { return arr[j+i] >= base }) + i
			out += (pFreq[n] - pFreq[j]) * freq[num]
			i = j
		}
	}
	return out % (1e9 + 7)
}
