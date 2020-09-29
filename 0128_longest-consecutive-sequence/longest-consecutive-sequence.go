package main

/*
  给定一个未排序的整数数组，找出最长连续序列的长度。
  要求算法的时间复杂度为 O(n)。
*/

// Hash
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
		count := 1
		for i := num + 1; numSet[i]; i++ {
			count++
		}
		out = Max(out, count)
	}
	return out
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
