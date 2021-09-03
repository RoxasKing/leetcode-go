package main

func arrayNesting(nums []int) int {
	n := len(nums)
	vis := make([]bool, n)
	out := 0
	for i := range nums {
		cnt := 0
		for x := i; !vis[x]; x = nums[x] {
			vis[x] = true
			cnt++
		}
		if cnt > out {
			out = cnt
		}
	}
	return out
}
