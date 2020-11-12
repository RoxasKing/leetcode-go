package main

/*
  给定长度分别为 m 和 n 的两个数组，其元素由 0-9 构成，表示两个自然数各位上的数字。现在从这两个数组中选出 k (k <= m + n) 个数字拼接成一个新的数，要求从同一个数组中取出的数字保持其在原数组中的相对顺序。
  求满足该条件的最大数。结果返回一个表示该最大数的长度为 k 的数组。

  说明: 请尽可能地优化你算法的时间和空间复杂度。

  示例 1:
    输入:
    nums1 = [3, 4, 6, 5]
    nums2 = [9, 1, 2, 5, 8, 3]
    k = 5
    输出:
    [9, 8, 6, 5, 3]

  示例 2:
    输入:
    nums1 = [6, 7]
    nums2 = [6, 0, 4]
    k = 5
    输出:
    [6, 7, 6, 0, 4]

  示例 3:
    输入:
    nums1 = [3, 9]
    nums2 = [8, 9]
    k = 3
    输出:
    [9, 8, 9]

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/create-maximum-number
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Greedy algorithm + Monotone Stack
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	out := []int{}
	n1, n2 := len(nums1), len(nums2)
	for i := 0; i <= k; i++ {
		if i > n1 || k-i > n2 {
			continue
		}
		tmp := merge(chooseMax(nums1, i), chooseMax(nums2, (k-i)))
		if isBigger(tmp, out) {
			out = tmp
		}
	}
	return out
}

func chooseMax(nums []int, k int) []int {
	del := len(nums) - k
	out := []int{}
	for _, num := range nums {
		for len(out) > 0 && out[len(out)-1] < num && del > 0 {
			out = out[:len(out)-1]
			del--
		}
		out = append(out, num)
	}
	return out[:k]
}

func merge(nums1, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	out := make([]int, n1+n2)
	i, j, k := 0, 0, 0
	for i < n1 && j < n2 {
		if isBigger(nums1[i:], nums2[j:]) {
			out[k] = nums1[i]
			i++
		} else {
			out[k] = nums2[j]
			j++
		}
		k++
	}
	copy(out[k:], nums1[i:])
	copy(out[k:], nums2[j:])
	return out
}

func isBigger(nums1, nums2 []int) bool {
	n1, n2 := len(nums1), len(nums2)
	for i := 0; i < Min(n1, n2); i++ {
		if nums1[i] != nums2[i] {
			return nums1[i] > nums2[i]
		}
	}
	return n1 > n2
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
