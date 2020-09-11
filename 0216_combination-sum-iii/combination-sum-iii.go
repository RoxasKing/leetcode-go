package main

/*
  找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。

  说明：
    所有数字都是正整数。
    解集不能包含重复的组合。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/combination-sum-iii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS + Backtracking
func combinationSum3(k int, n int) [][]int {
	var out [][]int
	var cur []int
	var curSum int
	var backTrack func(int)
	backTrack = func(index int) {
		if len(cur) == k {
			if curSum == n {
				tmp := make([]int, k)
				copy(tmp, cur)
				out = append(out, tmp)
			}
			return
		}
		if index == 10 {
			return
		}
		for i := index; i <= 9; i++ {
			cur = append(cur, i)
			curSum += i
			backTrack(i + 1)
			cur = cur[:len(cur)-1]
			curSum -= i
		}
	}
	backTrack(1)
	return out
}
