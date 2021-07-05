package main

// Tags:
// Backtracking
func subsets(nums []int) [][]int {
	out := [][]int{}
	dfs(nums, len(nums), 0, []int{}, &out)
	return out
}

func dfs(nums []int, n, i int, cur []int, out *[][]int) {
	if i == n {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*out = append(*out, tmp)
		return
	}

	cur = append(cur, nums[i])
	dfs(nums, n, i+1, cur, out)
	cur = cur[:len(cur)-1]

	dfs(nums, n, i+1, cur, out)
}
