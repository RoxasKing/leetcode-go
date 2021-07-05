package main

func monotoneIncreasingDigits(N int) int {
	if N < 10 {
		return N
	}
	nums := []int{}
	for n := N; n > 0; n /= 10 {
		nums = append(nums, n%10)
	}
	for i := 0; i < len(nums)>>1; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
	index := 0
	for index < len(nums)-1 && nums[index] <= nums[index+1] {
		index++
	}
	if index == len(nums)-1 {
		return N
	}
	nums[index]--
	for j := index + 1; j < len(nums); j++ {
		nums[j] = 9
	}
	for k := index - 1; k >= 0 && nums[k] > nums[k+1]; k-- {
		nums[k]--
		nums[k+1] = 9
	}
	if nums[0] == 0 {
		nums = nums[1:]
	}
	out := 0
	for _, num := range nums {
		out = out*10 + num
	}
	return out
}
