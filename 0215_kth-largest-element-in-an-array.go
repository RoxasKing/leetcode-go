package leetcode

import "math/rand"

/*
  在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

  说明:
    你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
*/

// Quick Sort
func findKthLargest(nums []int, k int) int {
	var quickSort func(int, int) int
	quickSort = func(l, r int) int {
		if l == r {
			return nums[l]
		}
		pivortIndex := l + rand.Intn(r+1-l)
		nums[pivortIndex], nums[r] = nums[r], nums[pivortIndex]
		index := l
		for i := l; i < r; i++ {
			if nums[i] > nums[r] {
				if nums[i] != nums[index] {
					nums[i], nums[index] = nums[index], nums[i]
				}
				index++
			}
		}
		nums[index], nums[r] = nums[r], nums[index]
		if k < index+1 {
			return quickSort(l, index)
		} else if k > index+1 {
			return quickSort(index+1, r)
		}
		return nums[index]
	}
	return quickSort(0, len(nums)-1)
}

// Heap Sort
func findKthLargest2(nums []int, k int) int {
	ajust := func() {
		for i := k>>1 - 1; i >= 0; i-- {
			son := i*2 + 1
			if son > k-1 {
				return
			}
			if son+1 < k && nums[son+1] < nums[son] {
				son++
			}
			if nums[son] < nums[i] {
				nums[son], nums[i] = nums[i], nums[son]
			}
		}
	}
	ajust()
	for i := k; i < len(nums); i++ {
		if nums[i] > nums[0] {
			nums[0], nums[i] = nums[i], nums[0]
			ajust()
		}
	}
	return nums[0]
}
