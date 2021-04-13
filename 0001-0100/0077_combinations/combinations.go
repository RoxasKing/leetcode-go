package main

/*
  给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

  示例:
    输入: n = 4, k = 2
    输出:
    [
      [2,4],
      [3,4],
      [2,3],
      [1,2],
      [1,3],
      [1,4],
    ]

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/combinations
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Backtracking
func combine(n int, k int) [][]int {
	cur := []int{}
	out := [][]int{}
	backtrack(n, k, 0, &cur, &out)
	return out
}

func backtrack(n int, k int, i int, cur *[]int, out *[][]int) {
	if len(*cur) == k {
		tmp := make([]int, k)
		copy(tmp, *cur)
		*out = append(*out, tmp)
		return
	}
	if i == n {
		return
	}
	*cur = append(*cur, i+1)
	backtrack(n, k, i+1, cur, out)
	*cur = (*cur)[:len(*cur)-1]

	backtrack(n, k, i+1, cur, out)
}
