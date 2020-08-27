package main

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

// Backtracking
func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	var out []string
	backtrack(&out, "", n, n)
	return out
}

func backtrack(out *[]string, cur string, lRemain, rRemain int) {
	if lRemain == 0 && rRemain == 0 {
		*out = append(*out, cur)
		return
	}
	if lRemain > 0 {
		backtrack(out, cur+"(", lRemain-1, rRemain)
	}
	if rRemain > 0 && rRemain > lRemain {
		backtrack(out, cur+")", lRemain, rRemain-1)
	}
}

// Iteration
func generateParenthesis2(n int) []string {
	dict := make([][]string, n+1)
	dict[0] = []string{""}
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			for _, s1 := range dict[j] {
				for _, s2 := range dict[i-1-j] {
					dict[i] = append(dict[i], "("+s1+")"+s2)
				}
			}
		}
	}
	return dict[n]
}
