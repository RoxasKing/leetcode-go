package main

/*
  在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。

  示例 1:
    输入: [7,5,6,4]
    输出: 5

  限制：
    0 <= 数组长度 <= 50000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Merge Sort
func reversePairs(nums []int) int {
	return merge(nums, 0, len(nums)-1)
}

func merge(nums []int, l, r int) int {
	if l >= r {
		return 0
	}
	m := l + (r-l)>>1
	cnt := merge(nums, l, m) + merge(nums, m+1, r)
	tmp := make([]int, r+1-l)
	i, j, k := l, m+1, 0
	for i <= m && j <= r {
		if nums[i] > nums[j] {
			tmp[k] = nums[j]
			j++
		} else {
			cnt += j - (m + 1)
			tmp[k] = nums[i]
			i++
		}
		k++
	}
	cnt += (m + 1 - i) * (r - m)
	copy(tmp[k:], nums[i:m+1])
	copy(tmp[k:], nums[j:r+1])
	copy(nums[l:r+1], tmp)
	return cnt
}
