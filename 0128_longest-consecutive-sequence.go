package leetcode

/*
  给定一个未排序的整数数组，找出最长连续序列的长度。
  要求算法的时间复杂度为 O(n)。
*/

func longestConsecutive(nums []int) int {
	numSet := make(map[int]struct{})
	for _, num := range nums {
		numSet[num] = struct{}{}
	}
	var out int
	for num := range numSet {
		if _, ok := numSet[num-1]; !ok {
			curNum := num
			cur := 1
			for {
				if _, ok := numSet[curNum+1]; ok {
					curNum++
					cur++
				} else {
					break
				}
			}
			if cur > out {
				out = cur
			}
		}
	}
	return out
}
