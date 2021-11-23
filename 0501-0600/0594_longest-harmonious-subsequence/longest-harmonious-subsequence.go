package main

// Difficulty:
// Easy

// Tags:
// Hash

func findLHS(nums []int) int {
	freq := map[int]int{}
	for _, num := range nums {
		freq[num]++
	}
	out := 0
	for num, c1 := range freq {
		if c2, ok := freq[num+1]; ok && c1+c2 > out {
			out = c1 + c2
		}
	}
	return out
}
