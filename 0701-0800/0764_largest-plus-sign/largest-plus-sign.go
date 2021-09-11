package main

// Tags:
// Dynamic Programming

func orderOfLargestPlusSign(n int, mines [][]int) int {
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	for _, x := range mines {
		f[x[0]][x[1]] = -1
	}

	for r := 0; r < n; r++ {
		cnt := 1
		for c := 0; c < n; c++ {
			if f[r][c] == -1 {
				cnt = 1
				continue
			}
			f[r][c] = cnt
			cnt++
		}
		cnt = 1
		for c := n - 1; c >= 0; c-- {
			if f[r][c] == -1 {
				cnt = 1
				continue
			}
			f[r][c] = Min(f[r][c], cnt)
			cnt++
		}
	}

	out := 0
	for c := 0; c < n; c++ {
		cnt := 1
		for r := 0; r < n; r++ {
			if f[r][c] == -1 {
				cnt = 1
				continue
			}
			f[r][c] = Min(f[r][c], cnt)
			cnt++
		}
		cnt = 1
		for r := n - 1; r >= 0; r-- {
			if f[r][c] == -1 {
				cnt = 1
				continue
			}
			f[r][c] = Min(f[r][c], cnt)
			cnt++
			out = Max(out, f[r][c])
		}
	}

	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
