package main

/*
  Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

  Example 1:
    Input: n = 3
    Output: ["((()))","(()())","(())()","()(())","()()()"]

  Example 2:
    Input: n = 1
    Output: ["()"]

  Constraints:
    1 <= n <= 8

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/generate-parentheses
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
