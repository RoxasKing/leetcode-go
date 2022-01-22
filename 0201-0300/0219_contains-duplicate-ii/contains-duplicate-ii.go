package main

// Difficulty:
// Easy

// Tags:
// Hash

func containsNearbyDuplicate(nums []int, k int) bool {
	dict := map[int]int{}
	for i, num := range nums {
		if j, ok := dict[num]; ok && i-j <= k {
			return true
		}
		dict[num] = i
	}
	return false
}
