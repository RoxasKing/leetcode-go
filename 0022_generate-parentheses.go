package leetcode

/*
  给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
  例如，给出 n = 3，生成结果为：
    [
      "((()))",
      "(()())",
      "(())()",
      "()(())",
      "()()()"
    ]
*/

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	var out []string
	dfs("", n, n, &out)
	return out
}

func dfs(str string, l, r int, out *[]string) {
	if l == 0 && r == 0 {
		*out = append(*out, str)
		return
	}
	if l > 0 {
		dfs(str+"(", l-1, r, out)
	}
	if r > 0 && r > l {
		dfs(str+")", l, r-1, out)
	}
}

func generateParenthesis2(n int) []string {
	var out []string
	switch n {
	case 0:
		out = append(out, "")
	default:
		for i := 0; i < n; i++ {
			for _, l := range generateParenthesis2(n - 1 - i) {
				for _, r := range generateParenthesis2(i) {
					out = append(out, "("+l+")"+r)
				}
			}
		}
	}
	return out
}
