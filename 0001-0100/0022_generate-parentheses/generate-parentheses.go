package main

// Tags:
// Backtracking
func generateParenthesis(n int) []string {
	out := []string{}
	dfs(n, n, "", &out)
	return out
}

func dfs(i, j int, cur string, out *[]string) {
	if i == 0 && j == 0 {
		*out = append(*out, cur)
		return
	}

	if i > 0 {
		dfs(i-1, j, cur+"(", out)
	}

	if j > i {
		dfs(i, j-1, cur+")", out)
	}
}
