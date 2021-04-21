package main

import (
	"sort"
)

/*
  Given a wooden stick of length n units. The stick is labelled from 0 to n. For example, a stick of length 6 is labelled as follows:

  Given an integer array cuts where cuts[i] denotes a position you should perform a cut at.

  You should perform the cuts in order, you can change the order of the cuts as you wish.

  The cost of one cut is the length of the stick to be cut, the total cost is the sum of costs of all cuts. When you cut a stick, it will be split into two smaller sticks (i.e. the sum of their lengths is the length of the stick before the cut). Please refer to the first example for a better explanation.

  Return the minimum total cost of the cuts.

  Example 1:
    Input: n = 7, cuts = [1,3,4,5]
    Output: 16
    Explanation: Using cuts order = [1, 3, 4, 5] as in the input leads to the following scenario:
      The first cut is done to a rod of length 7 so the cost is 7. The second cut is done to a rod of length 6 (i.e. the second part of the first cut), the third is done to a rod of length 4 and the last cut is to a rod of length 3. The total cost is 7 + 6 + 4 + 3 = 20.
      Rearranging the cuts to be [3, 5, 1, 4] for example will lead to a scenario with total cost = 16 (as shown in the example photo 7 + 4 + 3 + 2 = 16).

  Example 2:
    Input: n = 9, cuts = [5,6,1,4,2]
    Output: 22
    Explanation: If you try the given cuts ordering the cost will be 25.
    There are much ordering with total cost <= 25, for example, the order [4, 6, 5, 2, 1] has total cost = 22 which is the minimum possible.

  Constraints:
    1. 2 <= n <= 10^6
    2. 1 <= cuts.length <= min(n - 1, 100)
    3. 1 <= cuts[i] <= n - 1
    4. All the integers in cuts array are distinct.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minimum-cost-to-cut-a-stick
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func minCost(n int, cuts []int) int {
	sort.Ints(cuts)
	m := len(cuts) + 2
	lens := make([]int, m)
	copy(lens[1:m-1], cuts)
	lens[m-1] = n
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	for k := 2; k < m; k++ {
		for i := 0; i < m-k; i++ {
			dp[i][i+k] = 1<<31 - 1
			for j := 1; j < k; j++ {
				dp[i][i+k] = Min(dp[i][i+k], lens[i+k]-lens[i]+dp[i][i+j]+dp[i+j][i+k])
			}
		}
	}
	return dp[0][m-1]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
