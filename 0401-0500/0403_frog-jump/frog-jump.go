package main

import "sort"

/*
  A frog is crossing a river. The river is divided into some number of units, and at each unit, there may or may not exist a stone. The frog can jump on a stone, but it must not jump into the water.

  Given a list of stones' positions (in units) in sorted ascending order, determine if the frog can cross the river by landing on the last stone. Initially, the frog is on the first stone and assumes the first jump must be 1 unit.

  If the frog's last jump was k units, its next jump must be either k - 1, k, or k + 1 units. The frog can only jump in the forward direction.

  Example 1:
    Input: stones = [0,1,3,5,6,8,12,17]
    Output: true
    Explanation: The frog can jump to the last stone by jumping 1 unit to the 2nd stone, then 2 units to the 3rd stone, then 2 units to the 4th stone, then 3 units to the 6th stone, 4 units to the 7th stone, and 5 units to the 8th stone.

  Example 2:
    Input: stones = [0,1,2,3,4,8,9,11]
    Output: false
    Explanation: There is no way to jump to the last stone as the gap between the 5th and 6th stone is too large.


  Constraints:
    1. 2 <= stones.length <= 2000
    2. 0 <= stones[i] <= 2^31 - 1
    3. stones[0] == 0

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/frog-jump
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func canCross(stones []int) bool {
	dict := make(map[int]int)
	for i, stone := range stones {
		dict[stone] = i
	}
	n := len(stones)
	dp := make([][]int, n)
	dp[0] = append(dp[0], 0)
	for i := 0; i < n-1; i++ {
		sort.Ints(dp[i])
		prev := -1
		for _, step := range dp[i] {
			if step == prev {
				continue
			}
			prev = step
			if idx, ok := dict[stones[i]+step-1]; ok && stones[idx]-stones[i] == step-1 {
				dp[idx] = append(dp[idx], step-1)
			}
			if idx, ok := dict[stones[i]+step]; ok && stones[idx]-stones[i] == step {
				dp[idx] = append(dp[idx], step)
			}
			if idx, ok := dict[stones[i]+step+1]; ok && stones[idx]-stones[i] == step+1 {
				dp[idx] = append(dp[idx], step+1)
			}
		}
	}
	return len(dp[n-1]) > 0
}
