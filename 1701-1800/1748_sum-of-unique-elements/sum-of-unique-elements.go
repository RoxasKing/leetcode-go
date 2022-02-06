package main

// Difficulty:
// Easy

// Tags:
// Hash

func sumOfUnique(nums []int) int {
	freq := [101]int{}
	sum := 0
	for _, num := range nums {
		if freq[num]++; freq[num] == 1 {
			sum += num
		} else if freq[num] == 2 {
			sum -= num
		}
	}
	return sum
}
