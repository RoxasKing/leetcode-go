package main

import "math/rand"

// Difficulty:
// Medium

// Tags:
// Quick Select

func topKFrequent(nums []int, k int) []int {
	freq := map[int]int{}
	arr := []int{}
	for _, x := range nums {
		if _, ok := freq[x]; !ok {
			arr = append(arr, x)
		}
		freq[x]++
	}
	for l, r := 0, len(arr)-1; l < r; {
		rotateIdx := rand.Intn(r+1-l) + l
		arr[rotateIdx], arr[r] = arr[r], arr[rotateIdx]
		i, j := l-1, r
		for {
			for i++; i < r && freq[arr[i]] > freq[arr[r]]; i++ {
			}
			for j--; l < j && freq[arr[j]] < freq[arr[r]]; j-- {
			}
			if i >= j {
				break
			}
			arr[i], arr[j] = arr[j], arr[i]
		}
		arr[i], arr[r] = arr[r], arr[i]
		if i < k-1 {
			l = i + 1
		} else if i > k-1 {
			r = i - 1
		} else {
			break
		}
	}
	return arr[:k]
}
