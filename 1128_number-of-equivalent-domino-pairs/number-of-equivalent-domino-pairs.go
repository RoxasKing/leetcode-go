package main

/*
  Given a list of dominoes, dominoes[i] = [a, b] is equivalent to dominoes[j] = [c, d] if and only if either (a==c and b==d), or (a==d and b==c) - that is, one domino can be rotated to be equal to another domino.

  Return the number of pairs (i, j) for which 0 <= i < j < dominoes.length, and dominoes[i] is equivalent to dominoes[j].

  Example 1:
    Input: dominoes = [[1,2],[2,1],[3,4],[5,6]]
    Output: 1

  Constraints:
    1 <= dominoes.length <= 40000
    1 <= dominoes[i][j] <= 9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/number-of-equivalent-domino-pairs
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func numEquivDominoPairs(dominoes [][]int) int {
	count := make([]int, 100)
	for _, d := range dominoes {
		if d[0] > d[1] {
			d[0], d[1] = d[1], d[0]
		}
		count[d[0]*10+d[1]]++
	}
	out := 0
	for _, c := range count {
		out += c * (c - 1) / 2
	}
	return out
}
