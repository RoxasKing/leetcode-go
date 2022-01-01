package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Hash
// Sorting

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 || len(hand) < groupSize {
		return false
	}
	freq := map[int]int{}
	nums := []int{}
	for _, num := range hand {
		if _, ok := freq[num]; !ok {
			nums = append(nums, num)
		}
		freq[num]++
	}
	sort.Ints(nums)
	remain := len(hand)
	for i := 0; i <= len(nums)-groupSize; i++ {
		x := freq[nums[i]]
		if x == 0 {
			continue
		}
		for j := 0; j < groupSize; j++ {
			if v, ok := freq[nums[i]+j]; !ok || v < x {
				return false
			}
			freq[nums[i]+j] -= x
		}
		remain -= x * groupSize
	}
	return remain == 0
}
