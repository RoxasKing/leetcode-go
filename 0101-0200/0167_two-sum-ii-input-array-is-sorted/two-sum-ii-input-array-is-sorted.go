package main

// Tags:
// Two Pointers
func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	l, r := 0, n-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum < target {
			l++
		} else if sum > target {
			r--
		} else {
			break
		}
	}
	return []int{l + 1, r + 1}
}
