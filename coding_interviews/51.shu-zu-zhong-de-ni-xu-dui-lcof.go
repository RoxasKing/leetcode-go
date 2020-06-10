package codinginterviews

/*
  在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
*/

func reversePairs(nums []int) int {
	return mergeSortCount(nums, 0, len(nums)-1)
}

func mergeSortCount(nums []int, l, r int) int {
	if l >= r {
		return 0
	}
	m := l + (r-l)>>1
	count := mergeSortCount(nums, l, m) + mergeSortCount(nums, m+1, r)
	tmp := make([]int, r+1-l)
	var index int
	i, j := l, m+1
	for i <= m && j <= r {
		if nums[i] <= nums[j] {
			count += j - 1 - m
			tmp[index] = nums[i]
			index++
			i++
		} else {
			tmp[index] = nums[j]
			index++
			j++
		}
	}
	for i <= m {
		count += r - m
		tmp[index] = nums[i]
		index++
		i++
	}
	for j <= r {
		tmp[index] = nums[j]
		index++
		j++
	}
	copy(nums[l:r+1], tmp)
	return count
}
