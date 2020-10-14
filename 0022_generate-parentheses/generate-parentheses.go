package main

/*
  数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

  示例：
    输入：n = 3
    输出：[
           "((()))",
           "(()())",
           "(())()",
           "()(())",
           "()()()"
         ]

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/generate-parentheses
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
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
