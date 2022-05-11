package main

// Difficulty:
// Medium

// Tags:
// Backtracking

func combinationSum3(k int, n int) [][]int {
	o := [][]int{}
	t := []int{}
	x := 0
	var backtrack func(int)
	backtrack = func(i int) {
		if len(t) == k {
			if x == n {
				e := make([]int, k)
				copy(e, t)
				o = append(o, e)
			}
			return
		}
		if i == 10 {
			return
		}
		for ; i <= 9; i++ {
			t = append(t, i)
			x += i
			backtrack(i + 1)
			x -= i
			t = t[:len(t)-1]
		}
	}
	backtrack(1)
	return o
}
