package main

// Tags:
// Backtracking + DFS

func subsetXORSum(nums []int) int {
	out := 0
	dfs(nums, 0, 0, &out)
	return out
}

func dfs(nums []int, idx int, xor int, out *int) {
	if idx == len(nums) {
		*out += xor
		return
	}
	dfs(nums, idx+1, xor, out)
	dfs(nums, idx+1, xor^nums[idx], out)
}
