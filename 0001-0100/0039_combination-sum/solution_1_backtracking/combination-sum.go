package main

// Difficulty:
// Medium

// Tags:
// Backtracking

func combinationSum(candidates []int, target int) [][]int {
	n := len(candidates)
	out := [][]int{}
	cur := []int{}
	var dfs func(int)
	dfs = func(i int) {
		if target == 0 {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			out = append(out, tmp)
			return
		}
		if i == n {
			return
		}
		for k := 0; k*candidates[i] <= target; k++ {
			target -= k * candidates[i]
			for x := 0; x < k; x++ {
				cur = append(cur, candidates[i])
			}
			dfs(i + 1)
			cur = cur[:len(cur)-k]
			target += k * candidates[i]
		}
	}
	dfs(0)
	return out
}
