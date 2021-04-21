package main

/*
You have k lists of sorted integers in non-decreasing order. Find the smallest range that includes at least one number from each of the k lists.

We define the range [a, b] is smaller than range [c, d] if b - a < d - c or a < c if b - a == d - c.

Example 1:
  Input: nums = [[4,10,15,24,26],[0,9,12,20],[5,18,22,30]]
  Output: [20,24]
  Explanation:
  List 1: [4, 10, 15, 24,26], 24 is in range [20,24].
  List 2: [0, 9, 12, 20], 20 is in range [20,24].
  List 3: [5, 18, 22, 30], 22 is in range [20,24].

Example 2:
  Input: nums = [[1,2,3],[1,2,3],[1,2,3]]
  Output: [1,1]

Example 3:
  Input: nums = [[10,10],[11,11]]
  Output: [10,11]

Example 4:
  Input: nums = [[10],[11]]
  Output: [10,11]

Example 5:
  Input: nums = [[1],[2],[3],[4],[5],[6],[7]]
  Output: [1,7]

Constraints:
  nums.length == k
  1 <= k <= 3500
  1 <= nums[i].length <= 50
  -10^5 <= nums[i][j] <= 10^5
  nums[i] is sorted in non-decreasing order.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/smallest-range-covering-elements-from-k-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Two Pointers + Sliding Window
func smallestRange(nums [][]int) []int {
	mark := make(map[int][]int)
	min, max := 1<<31-1, -1<<31
	for i, arr := range nums {
		for _, num := range arr {
			mark[num] = append(mark[num], i)
			min = Min(min, num)
			max = Max(max, num)
		}
	}

	out := []int{min, max}
	cnt := make([]int, len(nums)) // 统计每个列表被包含的数
	inside := 0                   // 统计包含区间内数的列表个数
	for l, r := min, min; r <= max; r++ {
		if len(mark[r]) == 0 {
			continue
		}
		for _, idx := range mark[r] {
			if cnt[idx] == 0 {
				inside++
			}
			cnt[idx]++
		}
		for inside == len(nums) {
			if r-l < out[1]-out[0] {
				out = []int{l, r}
			}
			for _, idx := range mark[l] {
				cnt[idx]--
				if cnt[idx] == 0 {
					inside--
				}
			}
			l++
		}
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
