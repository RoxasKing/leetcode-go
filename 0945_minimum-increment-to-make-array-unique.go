package leetcode

import "sort"

/*
  给定整数数组 A，每次 move 操作将会选择任意 A[i]，并将其递增 1。
  返回使 A 中的每个值都是唯一的最少操作次数。
  示例 1:
    输入：[1,2,2]
    输出：1
    解释：经过一次 move 操作，数组将变为 [1, 2, 3]。
  示例 2:
    输入：[3,2,1,2,1,7]
    输出：6
    解释：经过 6 次 move 操作，数组将变为 [3, 4, 1, 2, 5, 7]。
    可以看出 5 次或 5 次以下的 move 操作是不能让数组的每个值唯一的。
  提示：
    0 <= A.length <= 40000
    0 <= A[i] < 40000
*/

func minIncrementForUnique(A []int) int {
	var count int
	sort.Ints(A)
	for i := 1; i < len(A); i++ {
		if A[i] <= A[i-1] {
			count += A[i-1] + 1 - A[i]
			A[i] = A[i-1] + 1
		}
	}
	return count
}

func minIncrementForUnique2(A []int) int {
	buckets := make([]int, 40001)
	max, min := 0, 40001
	for _, a := range A {
		buckets[a]++
		max = Max(max, a)
		min = Min(min, a)
	}
	var count, duplicate int
	for i := min; i < max; i++ {
		if buckets[i] > 1 {
			duplicate = buckets[i] - 1
			count += duplicate
			buckets[i] = 1
			buckets[i+1] += duplicate
		}
	}
	duplicate = buckets[max] - 1
	count += duplicate * (duplicate + 1) >> 1
	return count
}
