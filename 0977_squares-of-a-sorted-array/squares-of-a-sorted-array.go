package main

/*
  给定一个按非递减顺序排序的整数数组 A，返回每个数字的平方组成的新数组，要求也按非递减顺序排序。

  示例 1：
    输入：[-4,-1,0,3,10]
    输出：[0,1,9,16,100]

  示例 2：
    输入：[-7,-3,2,3,11]
    输出：[4,9,9,49,121]

  提示：
    1 <= A.length <= 10000
    -10000 <= A[i] <= 10000
    A 已按非递减顺序排序。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/squares-of-a-sorted-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func sortedSquares(A []int) []int {
	i := 0
	for i < len(A) && A[i] < 0 {
		i++
	}
	square(A)
	reverse(A[:i])
	return merge(A[:i], A[i:])
}

func reverse(nums []int) {
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func square(nums []int) {
	for i := range nums {
		nums[i] *= nums[i]
	}
}

func merge(A, B []int) []int {
	out := make([]int, 0, len(A)+len(B))
	for len(A) > 0 && len(B) > 0 {
		if A[0] < B[0] {
			out = append(out, A[0])
			A = A[1:]
		} else {
			out = append(out, B[0])
			B = B[1:]
		}
	}
	out = append(out, A...)
	out = append(out, B...)
	return out
}
