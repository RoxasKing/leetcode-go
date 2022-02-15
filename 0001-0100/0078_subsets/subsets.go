package main

// Difficulty:
// Medium

// Tags:
// Backtracking

func subsets(nums []int) [][]int {
	n := len(nums)
	out := [][]int{}
	cur := []int{}
	var k int
	var backtrack func(int)
	backtrack = func(i int) {
		if len(cur) == k {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			out = append(out, tmp)
			return
		}
		if i == n {
			return
		}
		cur = append(cur, nums[i])
		backtrack(i + 1)
		cur = cur[:len(cur)-1]
		backtrack(i + 1)
	}
	for k = 0; k <= n; k++ {
		backtrack(0)
	}
	return out
}
