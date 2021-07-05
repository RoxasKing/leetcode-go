package main

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	m := len(languages)
	cnt := make([]int, n+1)
	mark := make(map[int]bool)

	for _, f := range friendships {
		u, v := f[0], f[1]
		can := make([]bool, n+1)
		skip := false
		for _, lan := range languages[u-1] {
			can[lan] = true
		}
		for _, lan := range languages[v-1] {
			if can[lan] {
				skip = true
				break
			}
		}
		if skip {
			continue
		}

		if !mark[u] {
			for _, lan := range languages[u-1] {
				cnt[lan]++
			}
			mark[u] = true
		}
		if !mark[v] {
			for _, lan := range languages[v-1] {
				cnt[lan]++
			}
			mark[v] = true
		}
	}

	out := m
	for i := 1; i <= n; i++ {
		out = Min(out, len(mark)-cnt[i])
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
