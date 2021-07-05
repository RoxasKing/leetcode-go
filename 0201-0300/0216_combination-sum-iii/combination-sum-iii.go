package main

// Tags:
// Backtracking
func combinationSum3(k int, n int) [][]int {
	var out [][]int
	var cur []int
	var curSum int
	var backtrack func(int)
	backtrack = func(index int) {
		if len(cur) == k {
			if curSum == n {
				tmp := make([]int, k)
				copy(tmp, cur)
				out = append(out, tmp)
			}
			return
		}
		if index == 10 {
			return
		}
		for i := index; i <= 9; i++ {
			cur = append(cur, i)
			curSum += i
			backtrack(i + 1)
			cur = cur[:len(cur)-1]
			curSum -= i
		}
	}
	backtrack(1)
	return out
}
