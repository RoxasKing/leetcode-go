package main

// Tags:
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
