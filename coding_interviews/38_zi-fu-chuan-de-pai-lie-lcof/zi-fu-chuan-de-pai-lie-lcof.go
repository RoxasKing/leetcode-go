package main

// Tags:
// Backtracking
func permutation(s string) []string {
	out := []string{}
	chs := []byte(s)
	dfs(chs, 0, &out)
	return out
}

func dfs(chs []byte, i int, out *[]string) {
	if i == len(chs) {
		*out = append(*out, string(chs))
		return
	}
	mark := [128]bool{}
	for j := i; j < len(chs); j++ {
		if mark[chs[j]] {
			continue
		}
		mark[chs[j]] = true
		chs[i], chs[j] = chs[j], chs[i]
		dfs(chs, i+1, out)
		chs[i], chs[j] = chs[j], chs[i]
	}
}
