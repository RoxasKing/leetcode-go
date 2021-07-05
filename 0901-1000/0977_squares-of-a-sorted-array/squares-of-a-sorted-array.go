package main

// Tags:
// Merge Sort
func sortedSquares(nums []int) []int {
	n := len(nums)
	m := 0
	for m < n && nums[m] < 0 {
		m++
	}
	reverse(nums[:m])

	for i, num := range nums {
		nums[i] = num * num
	}

	out := make([]int, n)
	i, j, k := 0, m, 0
	for ; i < m && j < n; k++ {
		if nums[i] < nums[j] {
			out[k] = nums[i]
			i++
		} else {
			out[k] = nums[j]
			j++
		}
	}
	copy(out[k:], nums[i:m])
	copy(out[k:], nums[j:n])
	return out
}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n>>1; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
