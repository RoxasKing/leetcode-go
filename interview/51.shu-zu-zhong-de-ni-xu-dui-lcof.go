package interview

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
	var tmp []int
	i, j := l, m+1
	for i <= m && j <= r {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			count += j - (m + 1)
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for i <= m {
		tmp = append(tmp, nums[i])
		count += r - (m + 1) + 1
		i++
	}
	for j <= r {
		tmp = append(tmp, nums[j])
		j++
	}
	for i := l; i <= r; i++ {
		nums[i] = tmp[i-l]
	}
	return count
}
