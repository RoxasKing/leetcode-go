package main

// Difficulty:
// Easy

// Tags:
// Hash

func countQuadruplets(nums []int) int {
	var out int
	freq := map[int]int{}
	for b := len(nums) - 3; b >= 1; b-- {
		for _, numsd := range nums[b+2:] {
			freq[numsd-nums[b+1]]++
		}
		for _, numsa := range nums[:b] {
			if sum := numsa + nums[b]; freq[sum] > 0 {
				out += freq[sum]
			}
		}
	}
	return out
}
