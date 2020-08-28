package main

/*
  给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
*/

func combine(n int, k int) [][]int {
	var out [][]int
	var recur func(int, []int)
	recur = func(index int, array []int) {
		if len(array) == k {
			out = append(out, append(make([]int, 0, k), array...))
			return
		}
		for i := index; i <= n; i++ {
			array = append(array, i)
			recur(i+1, array)
			array = array[:len(array)-1]
		}
	}
	recur(1, make([]int, 0, k))
	return out
}
