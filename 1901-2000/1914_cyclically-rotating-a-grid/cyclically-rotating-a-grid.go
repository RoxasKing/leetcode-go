package main

func rotateGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	arr := [][]int{}
	for l, r, u, d := 0, n-1, 0, m-1; l <= r && u <= d; {
		cur := []int{}
		for i := l; i <= r; i++ {
			cur = append(cur, grid[u][i])
		}
		u++
		for i := u; i <= d; i++ {
			cur = append(cur, grid[i][r])
		}
		r--
		for i := r; i >= l; i-- {
			cur = append(cur, grid[d][i])
		}
		d--
		for i := d; i >= u; i-- {
			cur = append(cur, grid[i][l])
		}
		l++
		arr = append(arr, cur)
	}
	for i := range arr {
		x := k % len(arr[i])
		arr[i] = append(arr[i][x:], arr[i][:x]...)
	}
	for i, l, r, u, d := 0, 0, n-1, 0, m-1; l <= r && u <= d; i++ {
		cur := arr[i]
		for i := l; i <= r; i++ {
			grid[u][i] = cur[0]
			cur = cur[1:]
		}
		u++
		for i := u; i <= d; i++ {
			grid[i][r] = cur[0]
			cur = cur[1:]
		}
		r--
		for i := r; i >= l; i-- {
			grid[d][i] = cur[0]
			cur = cur[1:]
		}
		d--
		for i := d; i >= u; i-- {
			grid[i][l] = cur[0]
			cur = cur[1:]
		}
		l++
	}
	return grid
}
