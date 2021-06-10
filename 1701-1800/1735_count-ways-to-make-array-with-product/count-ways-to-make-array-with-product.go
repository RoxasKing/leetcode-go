package main

/*
  You are given a 2D integer array, queries. For each queries[i], where queries[i] = [ni, ki], find the number of different ways you can place positive integers into an array of size ni such that the product of the integers is ki. As the number of ways may be too large, the answer to the ith query is the number of ways modulo 109 + 7.

  Return an integer array answer where answer.length == queries.length, and answer[i] is the answer to the ith query.

  Example 1:
    Input: queries = [[2,6],[5,1],[73,660]]
    Output: [4,1,50734910]
    Explanation: Each query is independent.
      [2,6]: There are 4 ways to fill an array of size 2 that multiply to 6: [1,6], [2,3], [3,2], [6,1].
      [5,1]: There is 1 way to fill an array of size 5 that multiply to 1: [1,1,1,1,1].
      [73,660]: There are 1050734917 ways to fill an array of size 73 that multiply to 660. 1050734917 modulo 109 + 7 = 50734910.

  Example 2:
    Input: queries = [[1,1],[2,2],[3,3],[4,4],[5,5]]
    Output: [1,2,3,10,5]

  Constraints:
    1. 1 <= queries.length <= 10^4
    2. 1 <= ni, ki <= 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/count-ways-to-make-array-with-product
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Math

func waysToFillArray(queries [][]int) []int {
	f := make([][]int, 10000+13)
	for i := range f {
		f[i] = make([]int, 13+1)
	}
	f[0][0] = 1
	for i := 1; i < 10000+13; i++ {
		f[i][0] = 1
		for j := 1; j <= i && j <= 13; j++ {
			f[i][j] = f[i-1][j-1] + f[i-1][j]
			f[i][j] %= mod
		}
	}

	out := make([]int, len(queries))

	for i, q := range queries {
		n, k := q[0], q[1]
		out[i] = 1
		for j := 2; j*j <= k; j++ {
			if k%j > 0 {
				continue
			}
			r := 0
			for ; k%j == 0; k /= j {
				r++
			}
			out[i] *= f[r+n-1][r]
			out[i] %= mod
		}
		if k > 1 {
			out[i] *= n
			out[i] %= mod
		}
	}

	return out
}

var mod = int(1e9 + 7)
