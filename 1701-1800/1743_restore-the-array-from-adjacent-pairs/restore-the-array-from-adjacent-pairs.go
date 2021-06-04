package main

/*
  There is an integer array nums that consists of n unique elements, but you have forgotten it. However, you do remember every pair of adjacent elements in nums.

  You are given a 2D integer array adjacentPairs of size n - 1 where each adjacentPairs[i] = [ui, vi] indicates that the elements ui and vi are adjacent in nums.

  It is guaranteed that every adjacent pair of elements nums[i] and nums[i+1] will exist in adjacentPairs, either as [nums[i], nums[i+1]] or [nums[i+1], nums[i]]. The pairs can appear in any order.

  Return the original array nums. If there are multiple solutions, return any of them.

  Example 1:
    Input: adjacentPairs = [[2,1],[3,4],[3,2]]
    Output: [1,2,3,4]
    Explanation: This array has all its adjacent pairs in adjacentPairs.
      Notice that adjacentPairs[i] may not be in left-to-right order.

  Example 2:
    Input: adjacentPairs = [[4,-2],[1,4],[-3,1]]
    Output: [-2,4,1,-3]
    Explanation: There can be negative numbers.
      Another solution is [-3,1,4,-2], which would also be accepted.

  Example 3:
    Input: adjacentPairs = [[100000,-100000]]
    Output: [100000,-100000]

  Constraints:
    1. nums.length == n
    2. adjacentPairs.length == n - 1
    3. adjacentPairs[i].length == 2
    4. 2 <= n <= 10^5
    5. -10^5 <= nums[i], ui, vi <= 10^5
    6. There exists some nums that has adjacentPairs as its pairs.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/restore-the-array-from-adjacent-pairs
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func restoreArray(adjacentPairs [][]int) []int {
	edges := map[int][]int{}
	for _, p := range adjacentPairs {
		u, v := p[0], p[1]
		edges[u] = append(edges[u], v)
		edges[v] = append(edges[v], u)
	}

	var u int
	for num, arr := range edges {
		if len(arr) == 1 {
			u = num
			break
		}
	}

	out := make([]int, 0, len(adjacentPairs)+1)

	for {
		out = append(out, u)
		for _, v := range edges[u] {
			if len(out) > 1 && out[len(out)-2] == v {
				continue
			}
			u = v
		}
		if u == out[len(out)-1] {
			break
		}
	}

	return out
}
