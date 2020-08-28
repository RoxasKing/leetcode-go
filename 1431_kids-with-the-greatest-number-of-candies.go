package main

/*
  给你一个数组 candies 和一个整数 extraCandies ，其中 candies[i] 代表第 i 个孩子拥有的糖果数目。

  对每一个孩子，检查是否存在一种方案，将额外的 extraCandies 个糖果分配给孩子们之后，此孩子有 最多 的糖果。注意，允许有多个孩子同时拥有 最多 的糖果数目。

  提示：
    2 <= candies.length <= 100
    1 <= candies[i] <= 100
    1 <= extraCandies <= 50

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/kids-with-the-greatest-number-of-candies
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := -1
	for _, candy := range candies {
		max = Max(max, candy)
	}
	out := make([]bool, len(candies))
	for i := range candies {
		if candies[i]+extraCandies >= max {
			out[i] = true
		}
	}
	return out
}
