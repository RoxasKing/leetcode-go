package main

// Difficulty:
// Medium

// Tags:
// Two Pointers

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		if sum := numbers[l] + numbers[r]; sum == target {
			break
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
	return []int{l + 1, r + 1}
}
