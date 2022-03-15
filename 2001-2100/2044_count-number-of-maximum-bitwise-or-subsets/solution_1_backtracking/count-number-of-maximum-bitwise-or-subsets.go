package main

// Difficulty:
// Medium

// Tags:
// Backtracking

func countMaxOrSubsets(nums []int) int {
	var out, target, cur int
	var backtrack func(int)
	backtrack = func(i int) {
		if i == len(nums) {
			if cur == target {
				out++
			}
			return
		}
		pre := cur
		cur |= nums[i]
		backtrack(i + 1)
		cur = pre
		backtrack(i + 1)
	}
	for _, num := range nums {
		target |= num
	}
	backtrack(0)
	return out
}
