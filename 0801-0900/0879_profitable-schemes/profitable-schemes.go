package main

/*
  There is a group of n members, and a list of various crimes they could commit. The ith crime generates a profit[i] and requires group[i] members to participate in it. If a member participates in one crime, that member can't participate in another crime.

  Let's call a profitable scheme any subset of these crimes that generates at least minProfit profit, and the total number of members participating in that subset of crimes is at most n.

  Return the number of schemes that can be chosen. Since the answer may be very large, return it modulo 109 + 7.

  Example 1:
    Input: n = 5, minProfit = 3, group = [2,2], profit = [2,3]
    Output: 2
    Explanation: To make a profit of at least 3, the group could either commit crimes 0 and 1, or just crime 1.
      In total, there are 2 schemes.

  Example 2:
    Input: n = 10, minProfit = 5, group = [2,3,5], profit = [6,7,8]
    Output: 7
    Explanation: To make a profit of at least 5, the group could commit any crimes, as long as they commit one.
      There are 7 possible schemes: (0), (1), (2), (0,1), (0,2), (1,2), and (0,1,2).

  Constraints:
    1. 1 <= n <= 100
    2. 0 <= minProfit <= 100
    3. 1 <= group.length <= 100
    4. 1 <= group[i] <= 100
    5. profit.length == group.length
    6. 0 <= profit[i] <= 100

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/profitable-schemes
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Dynamic Programming

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, minProfit+1)
		f[i][0] = 1
	}

	for i := range group {
		for j := n; j >= group[i]; j-- {
			for k := 0; k <= minProfit; k++ {
				if profit[i]+k >= minProfit {
					f[j][minProfit] += f[j-group[i]][k]
					f[j][minProfit] %= 1e9 + 7
				} else {
					f[j][profit[i]+k] += f[j-group[i]][k]
					f[j][profit[i]+k] %= 1e9 + 7
				}
			}
		}
	}
	return f[n][minProfit]
}
