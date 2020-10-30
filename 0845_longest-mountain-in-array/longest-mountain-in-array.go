package main

/*
  我们把数组 A 中符合下列属性的任意连续子数组 B 称为 “山脉”：
    B.length >= 3
    存在 0 < i < B.length - 1 使得 B[0] < B[1] < ... B[i-1] < B[i] > B[i+1] > ... > B[B.length - 1]
    （注意：B 可以是 A 的任意子数组，包括整个数组 A。）
  给出一个整数数组 A，返回最长 “山脉” 的长度。
  如果不含有 “山脉” 则返回 0。

  示例 1：
    输入：[2,1,4,7,3,2,5]
    输出：5
    解释：最长的 “山脉” 是 [1,4,7,3,2]，长度为 5。

  示例 2：
    输入：[2,2,2]
    输出：0
    解释：不含 “山脉”。

  提示：
    0 <= A.length <= 10000
    0 <= A[i] <= 10000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/longest-mountain-in-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func longestMountain(a []int) (ans int) {
	n := len(a)
	L := make([]int, n)
	for i := 1; i < n-1; i++ {
		if a[i] > a[i-1] {
			L[i] = L[i-1] + 1
		}
	}
	R := make([]int, n)
	for i := n - 2; i > 0; i-- {
		if a[i] > a[i+1] {
			R[i] = R[i+1] + 1
		}
	}
	for i := 1; i < n-1; i++ {
		if L[i] > 0 && R[i] > 0 {
			ans = Max(ans, L[i]+R[i]+1)
		}
	}
	return
}

func longestMountain2(a []int) (ans int) {
	n := len(a)
	for l, r := 0, 1; l < n-2; l, r = r, r+1 {
		if a[l] >= a[r] {
			continue
		}
		for r < n-1 && a[r] < a[r+1] {
			r++
		}
		if r == n-1 || a[r] == a[r+1] {
			r++
			continue
		}
		for r < n-1 && a[r] > a[r+1] {
			r++
		}
		ans = Max(ans, r+1-l)
	}
	return
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
