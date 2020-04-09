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

// BackTrack
func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	var out []string
	generateParenthesisBackTrack("", n, n, &out)
	return out
}

func generateParenthesisBackTrack(str string, l, r int, out *[]string) {
	if l == 0 && r == 0 {
		*out = append(*out, str)
		return
	}
	if l > 0 {
		generateParenthesisBackTrack(str+"(", l-1, r, out)
	}
	if r > 0 && r > l {
		generateParenthesisBackTrack(str+")", l, r-1, out)
	}
}

// Recursive
func generateParenthesis2(n int) []string {
	dict := make([][]string, n+1)
	dict[0] = []string{""}
	return generateParenthesisRecursive(n, &dict)
}

func generateParenthesisRecursive(n int, dict *[][]string) []string {
	if len((*dict)[n]) != 0 {
		return (*dict)[n]
	}
	for i := 0; i < n; i++ {
		for _, l := range generateParenthesisRecursive(n-1-i, dict) {
			for _, r := range generateParenthesisRecursive(i, dict) {
				(*dict)[n] = append((*dict)[n], "("+l+")"+r)
			}
		}
	}
	return (*dict)[n]
}
