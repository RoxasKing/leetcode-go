package main

import "fmt"

func combine(n int, k int) [][]int {
	out := make([][]int, 0)

	var dfs func(int, []int)
	dfs = func(first int, curr []int) {
		if len(curr) == k {
			temp := make([]int, k)
			copy(temp, curr)
			out = append(out, temp)
			return
		}

		for i := first; i <= n-(k-len(curr))+1; i++ {
			curr = append(curr, i)
			dfs(i+1, curr)
			curr = curr[:len(curr)-1]
		}
	}
	dfs(1, make([]int, 0))

	return out
}

func main() {
	n, k := 4, 2
	fmt.Println(combine(n, k))
}
