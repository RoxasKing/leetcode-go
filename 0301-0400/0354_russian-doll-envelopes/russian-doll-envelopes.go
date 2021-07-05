package main

import "sort"

// Binary Search
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][1] != envelopes[j][1] {
			return envelopes[i][1] < envelopes[j][1]
		}
		return envelopes[i][0] > envelopes[j][0]
	})

	// e.g.          [[12 2] [3 4] [13 3] [4 5] [5 6] [14 4] [15 5] [5 5]]
	// after sort => [[12 2] [13 3] [14 4] [3 4] [15 5] [5 5] [4 5] [5 6]]
	// every heights can only selete one, and the width must be strictly incremented,
	// so we can get arrays like that:
	//               [[12 2] [13 3] [14 4] [15 5]]
	//            or [[3 4] [5 5]]
	//            or [[3 4] [4 5] [5 6]]
	// select the longest one

	arr := []int{}
	for _, envelope := range envelopes {
		i := sort.Search(len(arr), func(i int) bool { return arr[i] >= envelope[0] })
		if i < len(arr) {
			arr[i] = envelope[0]
			continue
		}
		arr = append(arr, envelope[0])
	}
	return len(arr)
}
