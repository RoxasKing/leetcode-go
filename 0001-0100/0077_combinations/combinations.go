package main

// Tags:
// Backtracking
func combine(n int, k int) [][]int {
	cur := []int{}
	out := [][]int{}
	backtrack(n, k, 0, &cur, &out)
	return out
}

func backtrack(n int, k int, i int, cur *[]int, out *[][]int) {
	if len(*cur) == k {
		tmp := make([]int, k)
		copy(tmp, *cur)
		*out = append(*out, tmp)
		return
	}
	if i == n {
		return
	}
	*cur = append(*cur, i+1)
	backtrack(n, k, i+1, cur, out)
	*cur = (*cur)[:len(*cur)-1]

	backtrack(n, k, i+1, cur, out)
}
