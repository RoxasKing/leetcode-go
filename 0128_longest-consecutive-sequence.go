package leetcode

/*
  给定一个未排序的整数数组，找出最长连续序列的长度。
  要求算法的时间复杂度为 O(n)。
*/

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}
	var out int
	for num := range numSet {
		if numSet[num-1] {
			continue
		}
		cur, count := num, 1
		for numSet[cur+1] {
			cur++
			count++
		}
		out = Max(out, count)
	}
	return out
}
