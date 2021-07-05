package main

// Tags:
// Math

func getPermutation(n int, k int) string {
	nums := make([]byte, 0, n)
	for i := 1; i <= n; i++ {
		nums = append(nums, byte(i)+'0')
	}

	k--
	base := 1
	for i := 2; i < n; i++ {
		base *= i
	}

	chs := make([]byte, n)
	for i := 0; i < n-1; i++ {
		idx := k / base
		chs[i] = nums[idx]
		copy(nums[idx:], nums[idx+1:])
		k %= base
		base /= n - 1 - i
	}
	chs[n-1] = nums[0]

	return string(chs)
}
