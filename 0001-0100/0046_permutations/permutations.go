package main

// Tags:
// Backtracking
func permute(nums []int) [][]int {
	out := [][]int{}
	dfs(nums, len(nums), 0, &out)
	return out
}

func dfs(nums []int, n, i int, out *[][]int) {
	if i == n {
		tmp := make([]int, n)
		copy(tmp, nums)
		*out = append(*out, tmp)
		return
	}

	for j := i; j < n; j++ {
		nums[i], nums[j] = nums[j], nums[i]
		dfs(nums, n, i+1, out)
		nums[i], nums[j] = nums[j], nums[i]
	}
}
