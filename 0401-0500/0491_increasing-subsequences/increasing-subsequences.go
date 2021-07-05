package main

// Tags:
// Backtracking
func findSubsequences(nums []int) [][]int {
	var out [][]int
	var cur []int
	backtrack(nums, &out, &cur, 0)
	return out
}

func backtrack(nums []int, out *[][]int, cur *[]int, index int) {
	if index == len(nums) {
		if len(*cur) >= 2 {
			tmp := make([]int, len(*cur))
			copy(tmp, *cur)
			*out = append(*out, tmp)
		}
		return
	}
	if len(*cur) == 0 || nums[index] >= (*cur)[len(*cur)-1] {
		*cur = append(*cur, nums[index])
		backtrack(nums, out, cur, index+1)
		*cur = (*cur)[:len(*cur)-1]
	}
	if len(*cur) == 0 || nums[index] != (*cur)[len(*cur)-1] {
		backtrack(nums, out, cur, index+1)
	}
}
