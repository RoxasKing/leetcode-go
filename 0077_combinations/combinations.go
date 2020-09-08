package main

/*
  给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
*/

// Backtracking
func combine(n int, k int) [][]int {
	var out [][]int
	cur := make([]int, 0, k)
	var backTrack func(int)
	backTrack = func(index int) {
		if len(cur) == k {
			tmp := make([]int, k)
			copy(tmp, cur)
			out = append(out, tmp)
			return
		}
		for i := index; i < n; i++ {
			cur = append(cur, i+1)
			backTrack(i + 1)
			cur = cur[:len(cur)-1]
		}
	}
	backTrack(0)
	return out
}
