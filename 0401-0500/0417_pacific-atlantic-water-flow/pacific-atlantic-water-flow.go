package main

// Difficulty:
// Medium

// Tags:
// DFS
// Bit Manipulation

func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}
	var dfs func(int, int, int, int)
	dfs = func(t, v, i, j int) {
		if i < 0 || m-1 < i || j < 0 || n-1 < j || v > heights[i][j] || t&f[i][j] == t {
			return
		}
		f[i][j] |= t
		dfs(t, heights[i][j], i-1, j)
		dfs(t, heights[i][j], i+1, j)
		dfs(t, heights[i][j], i, j-1)
		dfs(t, heights[i][j], i, j+1)
	}
	for i := 0; i < m; i++ {
		dfs(0b01, 0, i, 0)
		dfs(0b10, 0, i, n-1)
	}
	for j := 0; j < n; j++ {
		dfs(0b01, 0, 0, j)
		dfs(0b10, 0, m-1, j)
	}
	o := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if f[i][j] == 0b11 {
				o = append(o, []int{i, j})
			}
		}
	}
	return o
}
